// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Data {

    struct Identitas {
        string alamat;
        uint nomor;
        string nama;
        string status;
    }

    Identitas[] public identitas;

    string[] public dokumens;
    address public admin;

    constructor() {
        admin = msg.sender;
    }

    function addIdentitas(string memory alamat, uint nomor, string memory nama, string memory status) public {
        require(msg.sender == admin);
        identitas.push(Identitas({
            alamat: alamat,
            nomor: nomor,
            nama: nama,
            status: status
        }));
    } 

    function addDokumen(string memory dokumen) public {
        require(msg.sender == admin);
        dokumens.push(dokumen);
    } 

    function getIdentitas() external view returns (Identitas[] memory) {
        return identitas;
    }

    function getDokumens() external view returns (string[] memory) {
        return dokumens;
    }
}