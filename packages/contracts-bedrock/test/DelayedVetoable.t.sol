// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { DelayedVetoable_Init } from "./CommonTest.t.sol";
import { DelayedVetoable } from "src/L1/DelayedVetoable.sol";

contract DelayedVetoable_Getters_Test is DelayedVetoable_Init {
    /// @dev The getters return the expected values when called by the zero address.
    function test_getters() external {
        vm.startPrank(address(0));
        assertEq(delayedVetoable.initiator(), initiator);
        assertEq(delayedVetoable.vetoer(), vetoer);
        assertEq(delayedVetoable.target(), target);
        assertEq(delayedVetoable.delay(), operatingDelay);
        assertEq(delayedVetoable.queuedAt(keccak256(abi.encode(0))), 0);
    }
}

contract DelayedVetoable_Getters_TestFail is DelayedVetoable_Init {
    /// @dev Check that getter calls from unauthorized entities will revert.
    function test_getters_notZeroAddress_reverts() external {
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        delayedVetoable.initiator();
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        delayedVetoable.vetoer();
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        delayedVetoable.target();
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        delayedVetoable.delay();
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        delayedVetoable.queuedAt(keccak256(abi.encode(0)));
    }
}

contract DelayedVetoable_HandleCall_Test is DelayedVetoable_Init {
    /// @dev A call can be initiated by the initiator.
    function testFuzz_handleCall_initiation_succeeds(bytes calldata data) external {
        assumeNoClash(data);
        vm.expectEmit(true, false, false, true, address(delayedVetoable));
        emit Initiated(keccak256(data), data);

        vm.prank(initiator);
        (bool success,) = address(delayedVetoable).call(data);
        assertTrue(success);
    }

    /// @dev The delay is inititially set to zero and the call is immediately forwarded.
    function testFuzz_handleCall_initialForwardingImmediately_succeeds(
        bytes calldata inData,
        bytes calldata outData
    )
        external
    {
        assumeNonzeroData(inData);
        assumeNoClash(inData);

        // Reset the delay to zero
        vm.store(address(delayedVetoable), bytes32(uint256(2)), bytes32(uint256(0)));

        vm.mockCall(target, inData, outData);
        vm.expectEmit(true, false, false, true, address(delayedVetoable));
        vm.expectCall({ callee: target, data: inData });
        emit Forwarded(keccak256(inData), inData);
        vm.prank(initiator);
        (bool success, bytes memory returnData) = address(delayedVetoable).call(inData);
        assertTrue(success);
        assertEq(returnData, outData);

        // Check that the callHash is not stored for future forwarding
        bytes32 callHash = keccak256(inData);
        vm.prank(address(0));
        assertEq(delayedVetoable.queuedAt(callHash), 0);
    }

    /// @dev Calls are not forwarded until the delay has passed.
    function testFuzz_handleCall_forwardingWithDelay_succeeds(bytes calldata data) external {
        assumeNonzeroData(data);
        assumeNoClash(data);

        vm.prank(initiator);
        (bool success,) = address(delayedVetoable).call(data);

        // Check that the call is in the _queuedAt mapping
        bytes32 callHash = keccak256(data);
        vm.prank(address(0));
        assertEq(delayedVetoable.queuedAt(callHash), block.timestamp);

        vm.warp(block.timestamp + operatingDelay);
        vm.expectEmit(true, false, false, true, address(delayedVetoable));
        emit Forwarded(keccak256(data), data);

        vm.expectCall({ callee: target, data: data });
        (success,) = address(delayedVetoable).call(data);
        assertTrue(success);
    }
}

contract DelayedVetoable_HandleCall_TestFail is DelayedVetoable_Init {
    /// @dev Only the initiator can initiate a call.
    function test_handleCall_unauthorizedInitiation_reverts() external {
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        (bool success,) = address(delayedVetoable).call(NON_ZERO_DATA);
        assertTrue(success);
    }

    /// @dev The call cannot be forwarded until the delay has passed.
    function testFuzz_handleCall_forwardingTooSoon_reverts(bytes calldata data) external {
        assumeNoClash(data);
        vm.prank(initiator);
        (bool success,) = address(delayedVetoable).call(data);

        vm.expectRevert(abi.encodeWithSelector(ForwardingEarly.selector));
        (success,) = address(delayedVetoable).call(data);
        success;
    }

    /// @dev The call cannot be forwarded a second time.
    function testFuzz_handleCall_forwardingTwice_reverts(bytes calldata data) external {
        assumeNoClash(data);

        // Initiate the call
        vm.prank(initiator);
        (bool success,) = address(delayedVetoable).call(data);
        assertTrue(success);

        vm.warp(block.timestamp + operatingDelay);
        vm.expectEmit(true, false, false, true, address(delayedVetoable));
        emit Forwarded(keccak256(data), data);

        // Forward the call
        vm.expectCall({ callee: target, data: data });
        (success,) = address(delayedVetoable).call(data);
        assertTrue(success);

        // Attempt to foward the same call again.
        vm.expectRevert(abi.encodeWithSelector(Unauthorized.selector, initiator, address(this)));
        (success,) = address(delayedVetoable).call(data);
        assertTrue(success);
    }

    /// @dev If the target reverts, it is bubbled up.
    function testFuzz_handleCall_forwardingTargetReverts_reverts(
        bytes calldata inData,
        bytes calldata outData
    )
        external
    {
        assumeNoClash(inData);

        // Initiate the call
        vm.prank(initiator);
        (bool success,) = address(delayedVetoable).call(inData);
        success;

        vm.warp(block.timestamp + operatingDelay);
        vm.expectEmit(true, false, false, true, address(delayedVetoable));
        emit Forwarded(keccak256(inData), inData);

        vm.mockCallRevert(target, inData, outData);

        // Forward the call
        vm.expectRevert(outData);
        (bool success2,) = address(delayedVetoable).call(inData);
        success2;
    }

    /// @dev A test documenting the single instance in which the contract is not 'transparent' to the initiator.
    function testFuzz_handleCall_queuedAtClash_reverts(bytes memory outData) external {
        // This will get us calldata with the same function selector as the queuedAt function, but
        // with the incorrect input data length.
        bytes memory inData = abi.encodePacked(keccak256("queuedAt(bytes32)"));

        // Reset the delay to zero
        vm.store(address(delayedVetoable), bytes32(uint256(0)), bytes32(uint256(0)));

        vm.prank(initiator);
        vm.expectRevert(outData);
        (bool success,) = address(delayedVetoable).call(inData);
        success;
    }
}
