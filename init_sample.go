package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
// ���[�U�[�̐ݒ�
var userA, userB, userC int //string
//var Aval, Bval, Cval int
var err error

//init����
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
   
   userA = 100
   userB = 200
   userC = 400
   
    //Aval = 100
    //Bval = 200
    //Cval = 400
    
    //���[�U�[�A�J�E���g�Ɗ���U��͕ʓr�Ǘ����H�H����̎d�l�Ȃ烆�[�U�[�����͂��Đ��l�ݒ肷���OK�ȋC������B
    
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }

    return nil, nil
}
