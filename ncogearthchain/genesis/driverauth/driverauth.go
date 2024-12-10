package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from ncogearthchain-sfc c1d33c81f74abf82c0e22807f16e609578e10ad8, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b50612c90806100206000396000f3fe608060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630aeeca001461011757806318f628d4146101525780631e702f83146101f4578063242a6e3f14610239578063267ab446146102c95780634ddaf8f2146103045780634feb92f31461035557806366e7ea0f14610438578063715018a61461049357806379bead38146104aa5780638da5cb5b146105055780638f32d59b1461055c578063a4066fbe1461058b578063b9cc6b1c146105d0578063c0c53b8b14610656578063d6a0c7af146106e7578063e08d7e6614610758578063ebdf104c146107de578063f2fde38b14610963578063fd1b6ec1146109b4575b600080fd5b34801561012357600080fd5b506101506004803603602081101561013a57600080fd5b8101908080359060200190929190505050610a25565b005b34801561015e57600080fd5b506101f2600480360361012081101561017657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610b4d565b005b34801561020057600080fd5b506102376004803603604081101561021757600080fd5b810190808035906020019092919080359060200190929190505050610d58565b005b34801561024557600080fd5b506102c76004803603604081101561025c57600080fd5b81019080803590602001909291908035906020019064010000000081111561028357600080fd5b82018360208201111561029557600080fd5b803590602001918460018302840111640100000000831117156102b757600080fd5b9091929391929390505050610ef8565b005b3480156102d557600080fd5b50610302600480360360208110156102ec57600080fd5b810190808035906020019092919050505061109f565b005b34801561031057600080fd5b506103536004803603602081101561032757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111c7565b005b34801561036157600080fd5b50610436600480360361010081101561037957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156103c057600080fd5b8201836020820111156103d257600080fd5b803590602001918460018302840111640100000000831117156103f457600080fd5b9091929391929390803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919050505061131b565b005b34801561044457600080fd5b506104916004803603604081101561045b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061154a565b005b34801561049f57600080fd5b506104a8611804565b005b3480156104b657600080fd5b50610503600480360360408110156104cd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611941565b005b34801561051157600080fd5b5061051a611a9e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561056857600080fd5b50610571611ac8565b604051808215151515815260200191505060405180910390f35b34801561059757600080fd5b506105ce600480360360408110156105ae57600080fd5b810190808035906020019092919080359060200190929190505050611b20565b005b3480156105dc57600080fd5b50610654600480360360208110156105f357600080fd5b810190808035906020019064010000000081111561061057600080fd5b82018360208201111561062257600080fd5b8035906020019184600183028401116401000000008311171561064457600080fd5b9091929391929390505050611c9a565b005b34801561066257600080fd5b506106e56004803603606081101561067957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611def565b005b3480156106f357600080fd5b506107566004803603604081101561070a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fc1565b005b34801561076457600080fd5b506107dc6004803603602081101561077b57600080fd5b810190808035906020019064010000000081111561079857600080fd5b8201836020820111156107aa57600080fd5b803590602001918460208302840111640100000000831117156107cc57600080fd5b909192939192939050505061214a565b005b3480156107ea57600080fd5b506109616004803603608081101561080157600080fd5b810190808035906020019064010000000081111561081e57600080fd5b82018360208201111561083057600080fd5b8035906020019184602083028401116401000000008311171561085257600080fd5b90919293919293908035906020019064010000000081111561087357600080fd5b82018360208201111561088557600080fd5b803590602001918460208302840111640100000000831117156108a757600080fd5b9091929391929390803590602001906401000000008111156108c857600080fd5b8201836020820111156108da57600080fd5b803590602001918460208302840111640100000000831117156108fc57600080fd5b90919293919293908035906020019064010000000081111561091d57600080fd5b82018360208201111561092f57600080fd5b8035906020019184602083028401116401000000008311171561095157600080fd5b9091929391929390505050612311565b005b34801561096f57600080fd5b506109b26004803603602081101561098657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612583565b005b3480156109c057600080fd5b50610a23600480360360408110156109d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061260b565b005b610a2d611ac8565b1515610aa1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630aeeca00826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b158015610b3257600080fd5b505af1158015610b46573d6000803e3d6000fd5b5050505050565b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c38576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e81526020017f747261637400000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318f628d48a8a8a8a8a8a8a8a8a6040518a63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050600060405180830381600087803b158015610d3557600080fd5b505af1158015610d49573d6000803e3d6000fd5b50505050505050505050505050565b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e43576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e81526020017f747261637400000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e702f8383836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b158015610edc57600080fd5b505af1158015610ef0573d6000803e3d6000fd5b505050505050565b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616c6c6572206973206e6f74207468652053464320636f6e7472616374000081525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663242a6e3f8484846040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180848152602001806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b15801561108257600080fd5b505af1158015611096573d6000803e3d6000fd5b50505050505050565b6110a7611ac8565b151561111b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663267ab446826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156111ac57600080fd5b505af11580156111c0573d6000803e3d6000fd5b5050505050565b6111cf611ac8565b1515611243576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da7fc24f826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561130057600080fd5b505af1158015611314573d6000803e3d6000fd5b5050505050565b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e81526020017f747261637400000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b15801561152757600080fd5b505af115801561153b573d6000803e3d6000fd5b50505050505050505050505050565b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561160f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616c6c6572206973206e6f74207468652053464320636f6e7472616374000081525060200191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415156116fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f726563697069656e74206973206e6f74207468652053464320636f6e7472616381526020017f740000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e30443bc83611763848673ffffffffffffffffffffffffffffffffffffffff163161282290919063ffffffff16565b6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156117e857600080fd5b505af11580156117fc573d6000803e3d6000fd5b505050505050565b61180c611ac8565b1515611880576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b611949611ac8565b15156119bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166379bead3883836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611a8257600080fd5b505af1158015611a96573d6000803e3d6000fd5b505050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611be5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616c6c6572206973206e6f74207468652053464320636f6e7472616374000081525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a4066fbe83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b158015611c7e57600080fd5b505af1158015611c92573d6000803e3d6000fd5b505050505050565b611ca2611ac8565b1515611d16576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b9cc6b1c83836040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b158015611dd357600080fd5b505af1158015611de7573d6000803e3d6000fd5b505050505050565b600060019054906101000a900460ff1680611e0e5750611e0d6128ac565b5b80611e2557506000809054906101000a900460ff16155b1515611ebf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b60008060019054906101000a900460ff161590508015611f0f576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b611f18826128c3565b82606760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083606660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015611fbb5760008060016101000a81548160ff0219169083151502179055505b50505050565b611fc9611ac8565b151561203d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6a0c7af83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561212e57600080fd5b505af1158015612142573d6000803e3d6000fd5b505050505050565b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612235576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e81526020017f747261637400000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e08d7e6683836040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156122f557600080fd5b505af1158015612309573d6000803e3d6000fd5b505050505050565b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156123fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e81526020017f747261637400000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528b8b82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038352898982818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252878782818152602001925060200280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b15801561256157600080fd5b505af1158015612575573d6000803e3d6000fd5b505050505050505050505050565b61258b611ac8565b15156125ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61260881612ac6565b50565b612613611ac8565b1515612687576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61269082612c51565b80156126a157506126a081612c51565b5b1515612715576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f74206120636f6e747261637400000000000000000000000000000000000081525060200191505060405180910390fd5b606760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6a0c7af83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561280657600080fd5b505af115801561281a573d6000803e3d6000fd5b505050505050565b60008082840190508381101515156128a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000803090506000813b9050600081149250505090565b600060019054906101000a900460ff16806128e257506128e16128ac565b5b806128f957506000809054906101000a900460ff16155b1515612993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b60008060019054906101000a900460ff1615905080156129e3576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b81603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38015612ac25760008060016101000a81548160ff0219169083151502179055505b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515612b91576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526020017f646472657373000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080823b90506000811191505091905056fea165627a7a72305820fdf4aba53d7c11d5810fd82643a4b33539897091502fdacc0bcbd16bdd266e070029")
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0xd100ae0000000000000000000000000000000000")
