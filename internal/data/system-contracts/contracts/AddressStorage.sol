// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract AddressStorage is Ownable {
    mapping(address => Index) internal map;
    address[] internal addressList;

    struct Index {
        uint256 id;
    }

    constructor(address[] memory _addressList) {
        for (uint256 i = 0; i < _addressList.length; i++) {
            addressList.push(_addressList[i]);
            map[_addressList[i]].id = addressList.length;
        }
    }

    function add(address _address) external onlyOwner() returns (bool) {
        if (map[_address].id != 0) {
            return false;
        }
        _add(_address);
        return true;
    }

    function mustAdd(address _address) public onlyOwner() {
        require(
            map[_address].id == 0,
            "The address has already been added to the storage, failed to add the address to the address storage."
        );
        _add(_address);
    }

    function _add(address _address) private {
        addressList.push(_address);
        map[_address].id = addressList.length;
    }

    function size() external view returns (uint256) {
        return uint256(addressList.length);
    }

    function contains(address _address) external view returns (bool) {
        return map[_address].id > 0;
    }

    function getAddresses() external view returns (address[] memory) {
        return addressList;
    }

    function remove(address _address) external onlyOwner() returns (bool) {
        uint256 _id = map[_address].id;
        if (_id == 0 || _id > addressList.length) {
            return false;
        }
        _remove(_address);
        return true;
    }

    function mustRemove(address _address) external onlyOwner() {
        uint256 _id = map[_address].id;
        require(
            _id != 0,
            "The address does not exist, failed to remove the address from the address storage."
        );
        require(
            _id <= addressList.length,
            "Invalid index value for the address storage, failed to remove the address from the address storage."
        );

        _remove(_address);
    }

    function _remove(address _address) private {
        Index memory index = map[_address];

        // Move an last element of array into the vacated key slot.
        uint256 lastListID = addressList.length - 1;
        map[addressList[lastListID]].id = index.id;
        addressList[index.id - 1] = addressList[lastListID];

        addressList.pop();
        delete map[_address];
    }
}
