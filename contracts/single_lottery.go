// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SingleLotteryMetaData contains all meta data concerning the SingleLottery contract.
var SingleLotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTickets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ownerFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winnerPrizePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_returnedPrizePercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerPrize\",\"type\":\"uint256\"}],\"name\":\"LotteryDrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TicketPurchased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TICKETS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_FEE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RETURNED_PRIZE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TICKET_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WINNER_PRIZE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buyTickets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_randomSeed\",\"type\":\"uint256\"}],\"name\":\"drawLottery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMyInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tickets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainingTickets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lotteryFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"participantTickets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketsSold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040525f6001555f60025f6101000a81548160ff02191690831515021790555034801561002d575f5ffd5b50604051612772380380612772833981810160405281019061004f91906101d5565b5f8511610091576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610088906102cc565b60405180910390fd5b5f84116100d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ca9061035a565b60405180910390fd5b60648183856100e291906103a5565b6100ec91906103a5565b1461012c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012390610448565b60405180910390fd5b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084608081815250508360a081815250508260c081815250508160e081815250508061010081815250505050505050610466565b5f5ffd5b5f819050919050565b6101b4816101a2565b81146101be575f5ffd5b50565b5f815190506101cf816101ab565b92915050565b5f5f5f5f5f60a086880312156101ee576101ed61019e565b5b5f6101fb888289016101c1565b955050602061020c888289016101c1565b945050604061021d888289016101c1565b935050606061022e888289016101c1565b925050608061023f888289016101c1565b9150509295509295909350565b5f82825260208201905092915050565b7f5469636b6574207072696365206d7573742062652067726561746572207468615f8201527f6e207a65726f0000000000000000000000000000000000000000000000000000602082015250565b5f6102b660268361024c565b91506102c18261025c565b604082019050919050565b5f6020820190508181035f8301526102e3816102aa565b9050919050565b7f4d6178207469636b657473206d7573742062652067726561746572207468616e5f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f61034460258361024c565b915061034f826102ea565b604082019050919050565b5f6020820190508181035f83015261037181610338565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6103af826101a2565b91506103ba836101a2565b92508282019050808211156103d2576103d1610378565b5b92915050565b7f5072697a6520646973747269627574696f6e206d7573742073756d20746f20315f8201527f3030250000000000000000000000000000000000000000000000000000000000602082015250565b5f61043260238361024c565b915061043d826103d8565b604082019050919050565b5f6020820190508181035f83015261045f81610426565b9050919050565b60805160a05160c05160e051610100516122906104e25f395f818161052c015261129401525f8181610c35015261125501525f8181610d73015261121b01525f818161083d015281816108e301528181610c76015261174d01525f81816106f3015281816107c40152818161095301526110d001526122905ff3fe60806040526004361061013f575f3560e01c806370cc89ad116100b5578063b88a802f1161006e578063b88a802f14610434578063dfbf53ae1461044a578063e303a92b14610474578063eb2dbcf61461049c578063f2fde38b146104d8578063fca846dc146105005761013f565b806370cc89ad146103255780637daa10ce1461034f5780638da5cb5b1461037a5780638f15024f146103a4578063929066f5146103ce578063a8a3221e1461040a5761013f565b80632f366637116101075780632f3666371461021357806331d7a2621461022f57806335c1d3491461026b578063454055c0146102a757806361a25f07146102d15780636f9fb98a146102fb5761013f565b806302abf5fe14610143578063061fa2f91461016d57806316bfe25c146101a95780631a95f15f146101bf5780632017aa2f146101e9575b5f5ffd5b34801561014e575f5ffd5b5061015761052a565b6040516101649190611787565b60405180910390f35b348015610178575f5ffd5b50610193600480360381019061018e91906117fe565b61054e565b6040516101a09190611787565b60405180910390f35b3480156101b4575f5ffd5b506101bd610563565b005b3480156101ca575f5ffd5b506101d36107c2565b6040516101e09190611787565b60405180910390f35b3480156101f4575f5ffd5b506101fd6107e6565b60405161020a9190611787565b60405180910390f35b61022d60048036038101906102289190611853565b6107ec565b005b34801561023a575f5ffd5b50610255600480360381019061025091906117fe565b610be3565b6040516102629190611787565b60405180910390f35b348015610276575f5ffd5b50610291600480360381019061028c9190611853565b610bf8565b60405161029e919061188d565b60405180910390f35b3480156102b2575f5ffd5b506102bb610c33565b6040516102c89190611787565b60405180910390f35b3480156102dc575f5ffd5b506102e5610c57565b6040516102f291906118c0565b60405180910390f35b348015610306575f5ffd5b5061030f610c69565b60405161031c9190611787565b60405180910390f35b348015610330575f5ffd5b50610339610c70565b6040516103469190611787565b60405180910390f35b34801561035a575f5ffd5b50610363610ca4565b6040516103719291906118d9565b60405180910390f35b348015610385575f5ffd5b5061038e610d2a565b60405161039b919061188d565b60405180910390f35b3480156103af575f5ffd5b506103b8610d4e565b6040516103c59190611787565b60405180910390f35b3480156103d9575f5ffd5b506103f460048036038101906103ef91906117fe565b610d54565b60405161040191906118c0565b60405180910390f35b348015610415575f5ffd5b5061041e610d71565b60405161042b9190611787565b60405180910390f35b34801561043f575f5ffd5b50610448610d95565b005b348015610455575f5ffd5b5061045e610f3f565b60405161046b919061188d565b60405180910390f35b34801561047f575f5ffd5b5061049a60048036038101906104959190611853565b610f64565b005b3480156104a7575f5ffd5b506104c260048036038101906104bd9190611853565b611558565b6040516104cf919061188d565b60405180910390f35b3480156104e3575f5ffd5b506104fe60048036038101906104f991906117fe565b611593565b005b34801561050b575f5ffd5b5061051461174b565b6040516105219190611787565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b6006602052805f5260405f205f915090505481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e890611980565b60405180910390fd5b60025f9054906101000a900460ff1615610640576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610637906119e8565b60405180910390fd5b600160025f6101000a81548160ff0219169083151502179055505f5f90505b6003805490508110156107bf575f6003828154811061068157610680611a06565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205411156107b1575f7f000000000000000000000000000000000000000000000000000000000000000060065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461075a9190611a60565b90508060075f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107a89190611aa1565b92505081905550505b50808060010191505061065f565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b60095481565b60025f9054906101000a900460ff161561083b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083290611b1e565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001541061089f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089690611b86565b60405180910390fd5b5f81116108e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d890611bee565b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000816001546109109190611aa1565b1115610951576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094890611c56565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000008161097d9190611a60565b34146109be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b590611cbe565b60405180910390fd5b8060065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a0a9190611aa1565b9250508190555060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610b1557600333908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505b5f5f90505b81811015610ba657600533908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060015f815480929190610b9490611cdc565b91905055508080600101915050610b1a565b507f0668f5b446eb814fe35b3206f43f14bd8567ba04ddaf7a3ee56516929ab22ccb3382604051610bd8929190611d23565b60405180910390a150565b6007602052805f5260405f205f915090505481565b60038181548110610c07575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025f9054906101000a900460ff1681565b5f47905090565b5f6001547f0000000000000000000000000000000000000000000000000000000000000000610c9f9190611d4a565b905090565b5f5f60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205460075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054915091509091565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6004602052805f5260405f205f915054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f60075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8111610e18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0f90611dc7565b60405180910390fd5b5f60075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f3373ffffffffffffffffffffffffffffffffffffffff1682604051610e7f90611e12565b5f6040518083038185875af1925050503d805f8114610eb9576040519150601f19603f3d011682016040523d82523d5f602084013e610ebe565b606091505b5050905080610f02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef990611e70565b60405180910390fd5b7f106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f72413383604051610f33929190611d23565b60405180910390a15050565b60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ff2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe990611980565b60405180910390fd5b6002600380549050101561103b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103290611efe565b60405180910390fd5b5f6001541161107f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107690611f66565b60405180910390fd5b60025f9054906101000a900460ff16156110ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c590611b1e565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001546110fc9190611a60565b471461113d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113490611fce565b60405180910390fd5b600160025f6101000a81548160ff0219169083151502179055505f6001544244600585604051602001611173949392919061211c565b604051602081830303815290604052805190602001205f1c6111959190612192565b9050600581815481106111ab576111aa611a06565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f4790505f60647f0000000000000000000000000000000000000000000000000000000000000000836112459190611a60565b61124f91906121c2565b905060647f00000000000000000000000000000000000000000000000000000000000000008361127f9190611a60565b61128991906121c2565b6009819055505f60647f0000000000000000000000000000000000000000000000000000000000000000846112be9190611a60565b6112c891906121c2565b90508160075f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060095460075f60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461139b9190611aa1565b925050819055505f5f90505b6003805490508110156114f4575f600382815481106113c9576113c8611a06565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205411156114e6575f60015460065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054856114859190611a60565b61148f91906121c2565b90508060075f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114dd9190611aa1565b92505081905550505b5080806001019150506113a7565b507fe270ee2f324e081c14957e0abdb70aa14e5aacd0466ac05a9a1e89982cdb532660085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600954604051611549929190611d23565b60405180910390a15050505050565b60058181548110611567575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611621576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161161890611980565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361168f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116869061223c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f819050919050565b6117818161176f565b82525050565b5f60208201905061179a5f830184611778565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6117cd826117a4565b9050919050565b6117dd816117c3565b81146117e7575f5ffd5b50565b5f813590506117f8816117d4565b92915050565b5f60208284031215611813576118126117a0565b5b5f611820848285016117ea565b91505092915050565b6118328161176f565b811461183c575f5ffd5b50565b5f8135905061184d81611829565b92915050565b5f60208284031215611868576118676117a0565b5b5f6118758482850161183f565b91505092915050565b611887816117c3565b82525050565b5f6020820190506118a05f83018461187e565b92915050565b5f8115159050919050565b6118ba816118a6565b82525050565b5f6020820190506118d35f8301846118b1565b92915050565b5f6040820190506118ec5f830185611778565b6118f96020830184611778565b9392505050565b5f82825260208201905092915050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f61196a602183611900565b915061197582611910565b604082019050919050565b5f6020820190508181035f8301526119978161195e565b9050919050565b7f4c6f747465727920616c72656164792066696e697368656400000000000000005f82015250565b5f6119d2601883611900565b91506119dd8261199e565b602082019050919050565b5f6020820190508181035f8301526119ff816119c6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a6a8261176f565b9150611a758361176f565b9250828202611a838161176f565b91508282048414831517611a9a57611a99611a33565b5b5092915050565b5f611aab8261176f565b9150611ab68361176f565b9250828201905080821115611ace57611acd611a33565b5b92915050565b7f4c6f747465727920697320616c72656164792066696e697368656400000000005f82015250565b5f611b08601b83611900565b9150611b1382611ad4565b602082019050919050565b5f6020820190508181035f830152611b3581611afc565b9050919050565b7f416c6c207469636b6574732068617665206265656e20736f6c640000000000005f82015250565b5f611b70601a83611900565b9150611b7b82611b3c565b602082019050919050565b5f6020820190508181035f830152611b9d81611b64565b9050919050565b7f596f75206d75737420627579206174206c65617374206f6e65207469636b65745f82015250565b5f611bd8602083611900565b9150611be382611ba4565b602082019050919050565b5f6020820190508181035f830152611c0581611bcc565b9050919050565b7f4e6f7420656e6f756768207469636b65747320617661696c61626c65000000005f82015250565b5f611c40601c83611900565b9150611c4b82611c0c565b602082019050919050565b5f6020820190508181035f830152611c6d81611c34565b9050919050565b7f496e636f72726563742045544820616d6f756e742073656e74000000000000005f82015250565b5f611ca8601983611900565b9150611cb382611c74565b602082019050919050565b5f6020820190508181035f830152611cd581611c9c565b9050919050565b5f611ce68261176f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d1857611d17611a33565b5b600182019050919050565b5f604082019050611d365f83018561187e565b611d436020830184611778565b9392505050565b5f611d548261176f565b9150611d5f8361176f565b9250828203905081811115611d7757611d76611a33565b5b92915050565b7f4e6f207265776172647320746f20636c61696d000000000000000000000000005f82015250565b5f611db1601383611900565b9150611dbc82611d7d565b602082019050919050565b5f6020820190508181035f830152611dde81611da5565b9050919050565b5f81905092915050565b50565b5f611dfd5f83611de5565b9150611e0882611def565b5f82019050919050565b5f611e1c82611df2565b9150819050919050565b7f4661696c656420746f2073656e642072657761726400000000000000000000005f82015250565b5f611e5a601583611900565b9150611e6582611e26565b602082019050919050565b5f6020820190508181035f830152611e8781611e4e565b9050919050565b7f4174206c656173742074776f207061727469636970616e7473207265717569725f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ee8602283611900565b9150611ef382611e8e565b604082019050919050565b5f6020820190508181035f830152611f1581611edc565b9050919050565b7f4e6f207469636b65747320736f6c6420796574000000000000000000000000005f82015250565b5f611f50601383611900565b9150611f5b82611f1c565b602082019050919050565b5f6020820190508181035f830152611f7d81611f44565b9050919050565b7f5072697a6520706f6f6c20696e636f72726563740000000000000000000000005f82015250565b5f611fb8601483611900565b9150611fc382611f84565b602082019050919050565b5f6020820190508181035f830152611fe581611fac565b9050919050565b5f819050919050565b6120066120018261176f565b611fec565b82525050565b5f81549050919050565b5f81905092915050565b5f819050815f5260205f209050919050565b61203b816117c3565b82525050565b5f61204c8383612032565b60208301905092915050565b5f815f1c9050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61209461208f83612058565b612063565b9050919050565b5f6120a68254612082565b9050919050565b5f600182019050919050565b5f6120c38261200c565b6120cd8185612016565b93506120d883612020565b805f5b8381101561210f576120ec8261209b565b6120f68882612041565b9750612101836120ad565b9250506001810190506120db565b5085935050505092915050565b5f6121278287611ff5565b6020820191506121378286611ff5565b60208201915061214782856120b9565b91506121538284611ff5565b60208201915081905095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61219c8261176f565b91506121a78361176f565b9250826121b7576121b6612165565b5b828206905092915050565b5f6121cc8261176f565b91506121d78361176f565b9250826121e7576121e6612165565b5b828204905092915050565b7f4e6577206f776e657220697320746865207a65726f20616464726573730000005f82015250565b5f612226601d83611900565b9150612231826121f2565b602082019050919050565b5f6020820190508181035f8301526122538161221a565b905091905056fea264697066735822122044a2ccf5d825dd4968b92d07459f7c0d0bd8ca4c6b825e5151c0840cd19be3fc64736f6c634300081d0033",
}

// SingleLotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use SingleLotteryMetaData.ABI instead.
var SingleLotteryABI = SingleLotteryMetaData.ABI

// SingleLotteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SingleLotteryMetaData.Bin instead.
var SingleLotteryBin = SingleLotteryMetaData.Bin

// DeploySingleLottery deploys a new Ethereum contract, binding an instance of SingleLottery to it.
func DeploySingleLottery(auth *bind.TransactOpts, backend bind.ContractBackend, _ticketPrice *big.Int, _maxTickets *big.Int, _ownerFeePercent *big.Int, _winnerPrizePercent *big.Int, _returnedPrizePercent *big.Int) (common.Address, *types.Transaction, *SingleLottery, error) {
	parsed, err := SingleLotteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SingleLotteryBin), backend, _ticketPrice, _maxTickets, _ownerFeePercent, _winnerPrizePercent, _returnedPrizePercent)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SingleLottery{SingleLotteryCaller: SingleLotteryCaller{contract: contract}, SingleLotteryTransactor: SingleLotteryTransactor{contract: contract}, SingleLotteryFilterer: SingleLotteryFilterer{contract: contract}}, nil
}

// SingleLottery is an auto generated Go binding around an Ethereum contract.
type SingleLottery struct {
	SingleLotteryCaller     // Read-only binding to the contract
	SingleLotteryTransactor // Write-only binding to the contract
	SingleLotteryFilterer   // Log filterer for contract events
}

// SingleLotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleLotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleLotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleLotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleLotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingleLotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleLotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleLotterySession struct {
	Contract     *SingleLottery    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingleLotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleLotteryCallerSession struct {
	Contract *SingleLotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SingleLotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleLotteryTransactorSession struct {
	Contract     *SingleLotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SingleLotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleLotteryRaw struct {
	Contract *SingleLottery // Generic contract binding to access the raw methods on
}

// SingleLotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleLotteryCallerRaw struct {
	Contract *SingleLotteryCaller // Generic read-only contract binding to access the raw methods on
}

// SingleLotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleLotteryTransactorRaw struct {
	Contract *SingleLotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleLottery creates a new instance of SingleLottery, bound to a specific deployed contract.
func NewSingleLottery(address common.Address, backend bind.ContractBackend) (*SingleLottery, error) {
	contract, err := bindSingleLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleLottery{SingleLotteryCaller: SingleLotteryCaller{contract: contract}, SingleLotteryTransactor: SingleLotteryTransactor{contract: contract}, SingleLotteryFilterer: SingleLotteryFilterer{contract: contract}}, nil
}

// NewSingleLotteryCaller creates a new read-only instance of SingleLottery, bound to a specific deployed contract.
func NewSingleLotteryCaller(address common.Address, caller bind.ContractCaller) (*SingleLotteryCaller, error) {
	contract, err := bindSingleLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleLotteryCaller{contract: contract}, nil
}

// NewSingleLotteryTransactor creates a new write-only instance of SingleLottery, bound to a specific deployed contract.
func NewSingleLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleLotteryTransactor, error) {
	contract, err := bindSingleLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleLotteryTransactor{contract: contract}, nil
}

// NewSingleLotteryFilterer creates a new log filterer instance of SingleLottery, bound to a specific deployed contract.
func NewSingleLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleLotteryFilterer, error) {
	contract, err := bindSingleLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleLotteryFilterer{contract: contract}, nil
}

// bindSingleLottery binds a generic wrapper to an already deployed contract.
func bindSingleLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SingleLotteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleLottery *SingleLotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleLottery.Contract.SingleLotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleLottery *SingleLotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleLottery.Contract.SingleLotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleLottery *SingleLotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleLottery.Contract.SingleLotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleLottery *SingleLotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleLottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleLottery *SingleLotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleLottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleLottery *SingleLotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleLottery.Contract.contract.Transact(opts, method, params...)
}

// MAXTICKETS is a free data retrieval call binding the contract method 0xfca846dc.
//
// Solidity: function MAX_TICKETS() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) MAXTICKETS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "MAX_TICKETS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTICKETS is a free data retrieval call binding the contract method 0xfca846dc.
//
// Solidity: function MAX_TICKETS() view returns(uint256)
func (_SingleLottery *SingleLotterySession) MAXTICKETS() (*big.Int, error) {
	return _SingleLottery.Contract.MAXTICKETS(&_SingleLottery.CallOpts)
}

// MAXTICKETS is a free data retrieval call binding the contract method 0xfca846dc.
//
// Solidity: function MAX_TICKETS() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) MAXTICKETS() (*big.Int, error) {
	return _SingleLottery.Contract.MAXTICKETS(&_SingleLottery.CallOpts)
}

// OWNERFEEPERCENT is a free data retrieval call binding the contract method 0xa8a3221e.
//
// Solidity: function OWNER_FEE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) OWNERFEEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "OWNER_FEE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OWNERFEEPERCENT is a free data retrieval call binding the contract method 0xa8a3221e.
//
// Solidity: function OWNER_FEE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotterySession) OWNERFEEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.OWNERFEEPERCENT(&_SingleLottery.CallOpts)
}

// OWNERFEEPERCENT is a free data retrieval call binding the contract method 0xa8a3221e.
//
// Solidity: function OWNER_FEE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) OWNERFEEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.OWNERFEEPERCENT(&_SingleLottery.CallOpts)
}

// RETURNEDPRIZEPERCENT is a free data retrieval call binding the contract method 0x02abf5fe.
//
// Solidity: function RETURNED_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) RETURNEDPRIZEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "RETURNED_PRIZE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RETURNEDPRIZEPERCENT is a free data retrieval call binding the contract method 0x02abf5fe.
//
// Solidity: function RETURNED_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotterySession) RETURNEDPRIZEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.RETURNEDPRIZEPERCENT(&_SingleLottery.CallOpts)
}

// RETURNEDPRIZEPERCENT is a free data retrieval call binding the contract method 0x02abf5fe.
//
// Solidity: function RETURNED_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) RETURNEDPRIZEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.RETURNEDPRIZEPERCENT(&_SingleLottery.CallOpts)
}

// TICKETPRICE is a free data retrieval call binding the contract method 0x1a95f15f.
//
// Solidity: function TICKET_PRICE() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) TICKETPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "TICKET_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TICKETPRICE is a free data retrieval call binding the contract method 0x1a95f15f.
//
// Solidity: function TICKET_PRICE() view returns(uint256)
func (_SingleLottery *SingleLotterySession) TICKETPRICE() (*big.Int, error) {
	return _SingleLottery.Contract.TICKETPRICE(&_SingleLottery.CallOpts)
}

// TICKETPRICE is a free data retrieval call binding the contract method 0x1a95f15f.
//
// Solidity: function TICKET_PRICE() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) TICKETPRICE() (*big.Int, error) {
	return _SingleLottery.Contract.TICKETPRICE(&_SingleLottery.CallOpts)
}

// WINNERPRIZEPERCENT is a free data retrieval call binding the contract method 0x454055c0.
//
// Solidity: function WINNER_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) WINNERPRIZEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "WINNER_PRIZE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WINNERPRIZEPERCENT is a free data retrieval call binding the contract method 0x454055c0.
//
// Solidity: function WINNER_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotterySession) WINNERPRIZEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.WINNERPRIZEPERCENT(&_SingleLottery.CallOpts)
}

// WINNERPRIZEPERCENT is a free data retrieval call binding the contract method 0x454055c0.
//
// Solidity: function WINNER_PRIZE_PERCENT() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) WINNERPRIZEPERCENT() (*big.Int, error) {
	return _SingleLottery.Contract.WINNERPRIZEPERCENT(&_SingleLottery.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_SingleLottery *SingleLotterySession) GetContractBalance() (*big.Int, error) {
	return _SingleLottery.Contract.GetContractBalance(&_SingleLottery.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) GetContractBalance() (*big.Int, error) {
	return _SingleLottery.Contract.GetContractBalance(&_SingleLottery.CallOpts)
}

// GetMyInfo is a free data retrieval call binding the contract method 0x7daa10ce.
//
// Solidity: function getMyInfo() view returns(uint256 tickets, uint256 rewards)
func (_SingleLottery *SingleLotteryCaller) GetMyInfo(opts *bind.CallOpts) (struct {
	Tickets *big.Int
	Rewards *big.Int
}, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "getMyInfo")

	outstruct := new(struct {
		Tickets *big.Int
		Rewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tickets = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMyInfo is a free data retrieval call binding the contract method 0x7daa10ce.
//
// Solidity: function getMyInfo() view returns(uint256 tickets, uint256 rewards)
func (_SingleLottery *SingleLotterySession) GetMyInfo() (struct {
	Tickets *big.Int
	Rewards *big.Int
}, error) {
	return _SingleLottery.Contract.GetMyInfo(&_SingleLottery.CallOpts)
}

// GetMyInfo is a free data retrieval call binding the contract method 0x7daa10ce.
//
// Solidity: function getMyInfo() view returns(uint256 tickets, uint256 rewards)
func (_SingleLottery *SingleLotteryCallerSession) GetMyInfo() (struct {
	Tickets *big.Int
	Rewards *big.Int
}, error) {
	return _SingleLottery.Contract.GetMyInfo(&_SingleLottery.CallOpts)
}

// GetRemainingTickets is a free data retrieval call binding the contract method 0x70cc89ad.
//
// Solidity: function getRemainingTickets() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) GetRemainingTickets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "getRemainingTickets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingTickets is a free data retrieval call binding the contract method 0x70cc89ad.
//
// Solidity: function getRemainingTickets() view returns(uint256)
func (_SingleLottery *SingleLotterySession) GetRemainingTickets() (*big.Int, error) {
	return _SingleLottery.Contract.GetRemainingTickets(&_SingleLottery.CallOpts)
}

// GetRemainingTickets is a free data retrieval call binding the contract method 0x70cc89ad.
//
// Solidity: function getRemainingTickets() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) GetRemainingTickets() (*big.Int, error) {
	return _SingleLottery.Contract.GetRemainingTickets(&_SingleLottery.CallOpts)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_SingleLottery *SingleLotteryCaller) IsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "isParticipant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_SingleLottery *SingleLotterySession) IsParticipant(arg0 common.Address) (bool, error) {
	return _SingleLottery.Contract.IsParticipant(&_SingleLottery.CallOpts, arg0)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_SingleLottery *SingleLotteryCallerSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _SingleLottery.Contract.IsParticipant(&_SingleLottery.CallOpts, arg0)
}

// LotteryFinished is a free data retrieval call binding the contract method 0x61a25f07.
//
// Solidity: function lotteryFinished() view returns(bool)
func (_SingleLottery *SingleLotteryCaller) LotteryFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "lotteryFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LotteryFinished is a free data retrieval call binding the contract method 0x61a25f07.
//
// Solidity: function lotteryFinished() view returns(bool)
func (_SingleLottery *SingleLotterySession) LotteryFinished() (bool, error) {
	return _SingleLottery.Contract.LotteryFinished(&_SingleLottery.CallOpts)
}

// LotteryFinished is a free data retrieval call binding the contract method 0x61a25f07.
//
// Solidity: function lotteryFinished() view returns(bool)
func (_SingleLottery *SingleLotteryCallerSession) LotteryFinished() (bool, error) {
	return _SingleLottery.Contract.LotteryFinished(&_SingleLottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleLottery *SingleLotteryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleLottery *SingleLotterySession) Owner() (common.Address, error) {
	return _SingleLottery.Contract.Owner(&_SingleLottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleLottery *SingleLotteryCallerSession) Owner() (common.Address, error) {
	return _SingleLottery.Contract.Owner(&_SingleLottery.CallOpts)
}

// ParticipantTickets is a free data retrieval call binding the contract method 0x061fa2f9.
//
// Solidity: function participantTickets(address ) view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) ParticipantTickets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "participantTickets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParticipantTickets is a free data retrieval call binding the contract method 0x061fa2f9.
//
// Solidity: function participantTickets(address ) view returns(uint256)
func (_SingleLottery *SingleLotterySession) ParticipantTickets(arg0 common.Address) (*big.Int, error) {
	return _SingleLottery.Contract.ParticipantTickets(&_SingleLottery.CallOpts, arg0)
}

// ParticipantTickets is a free data retrieval call binding the contract method 0x061fa2f9.
//
// Solidity: function participantTickets(address ) view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) ParticipantTickets(arg0 common.Address) (*big.Int, error) {
	return _SingleLottery.Contract.ParticipantTickets(&_SingleLottery.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_SingleLottery *SingleLotteryCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_SingleLottery *SingleLotterySession) Participants(arg0 *big.Int) (common.Address, error) {
	return _SingleLottery.Contract.Participants(&_SingleLottery.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_SingleLottery *SingleLotteryCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _SingleLottery.Contract.Participants(&_SingleLottery.CallOpts, arg0)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address ) view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) PendingRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "pendingRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address ) view returns(uint256)
func (_SingleLottery *SingleLotterySession) PendingRewards(arg0 common.Address) (*big.Int, error) {
	return _SingleLottery.Contract.PendingRewards(&_SingleLottery.CallOpts, arg0)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address ) view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) PendingRewards(arg0 common.Address) (*big.Int, error) {
	return _SingleLottery.Contract.PendingRewards(&_SingleLottery.CallOpts, arg0)
}

// TicketOwners is a free data retrieval call binding the contract method 0xeb2dbcf6.
//
// Solidity: function ticketOwners(uint256 ) view returns(address)
func (_SingleLottery *SingleLotteryCaller) TicketOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "ticketOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TicketOwners is a free data retrieval call binding the contract method 0xeb2dbcf6.
//
// Solidity: function ticketOwners(uint256 ) view returns(address)
func (_SingleLottery *SingleLotterySession) TicketOwners(arg0 *big.Int) (common.Address, error) {
	return _SingleLottery.Contract.TicketOwners(&_SingleLottery.CallOpts, arg0)
}

// TicketOwners is a free data retrieval call binding the contract method 0xeb2dbcf6.
//
// Solidity: function ticketOwners(uint256 ) view returns(address)
func (_SingleLottery *SingleLotteryCallerSession) TicketOwners(arg0 *big.Int) (common.Address, error) {
	return _SingleLottery.Contract.TicketOwners(&_SingleLottery.CallOpts, arg0)
}

// TicketsSold is a free data retrieval call binding the contract method 0x8f15024f.
//
// Solidity: function ticketsSold() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) TicketsSold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "ticketsSold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketsSold is a free data retrieval call binding the contract method 0x8f15024f.
//
// Solidity: function ticketsSold() view returns(uint256)
func (_SingleLottery *SingleLotterySession) TicketsSold() (*big.Int, error) {
	return _SingleLottery.Contract.TicketsSold(&_SingleLottery.CallOpts)
}

// TicketsSold is a free data retrieval call binding the contract method 0x8f15024f.
//
// Solidity: function ticketsSold() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) TicketsSold() (*big.Int, error) {
	return _SingleLottery.Contract.TicketsSold(&_SingleLottery.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_SingleLottery *SingleLotteryCaller) Winner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "winner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_SingleLottery *SingleLotterySession) Winner() (common.Address, error) {
	return _SingleLottery.Contract.Winner(&_SingleLottery.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_SingleLottery *SingleLotteryCallerSession) Winner() (common.Address, error) {
	return _SingleLottery.Contract.Winner(&_SingleLottery.CallOpts)
}

// WinnerPrize is a free data retrieval call binding the contract method 0x2017aa2f.
//
// Solidity: function winnerPrize() view returns(uint256)
func (_SingleLottery *SingleLotteryCaller) WinnerPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleLottery.contract.Call(opts, &out, "winnerPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinnerPrize is a free data retrieval call binding the contract method 0x2017aa2f.
//
// Solidity: function winnerPrize() view returns(uint256)
func (_SingleLottery *SingleLotterySession) WinnerPrize() (*big.Int, error) {
	return _SingleLottery.Contract.WinnerPrize(&_SingleLottery.CallOpts)
}

// WinnerPrize is a free data retrieval call binding the contract method 0x2017aa2f.
//
// Solidity: function winnerPrize() view returns(uint256)
func (_SingleLottery *SingleLotteryCallerSession) WinnerPrize() (*big.Int, error) {
	return _SingleLottery.Contract.WinnerPrize(&_SingleLottery.CallOpts)
}

// BuyTickets is a paid mutator transaction binding the contract method 0x2f366637.
//
// Solidity: function buyTickets(uint256 _amount) payable returns()
func (_SingleLottery *SingleLotteryTransactor) BuyTickets(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SingleLottery.contract.Transact(opts, "buyTickets", _amount)
}

// BuyTickets is a paid mutator transaction binding the contract method 0x2f366637.
//
// Solidity: function buyTickets(uint256 _amount) payable returns()
func (_SingleLottery *SingleLotterySession) BuyTickets(_amount *big.Int) (*types.Transaction, error) {
	return _SingleLottery.Contract.BuyTickets(&_SingleLottery.TransactOpts, _amount)
}

// BuyTickets is a paid mutator transaction binding the contract method 0x2f366637.
//
// Solidity: function buyTickets(uint256 _amount) payable returns()
func (_SingleLottery *SingleLotteryTransactorSession) BuyTickets(_amount *big.Int) (*types.Transaction, error) {
	return _SingleLottery.Contract.BuyTickets(&_SingleLottery.TransactOpts, _amount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_SingleLottery *SingleLotteryTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleLottery.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_SingleLottery *SingleLotterySession) ClaimReward() (*types.Transaction, error) {
	return _SingleLottery.Contract.ClaimReward(&_SingleLottery.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_SingleLottery *SingleLotteryTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _SingleLottery.Contract.ClaimReward(&_SingleLottery.TransactOpts)
}

// DrawLottery is a paid mutator transaction binding the contract method 0xe303a92b.
//
// Solidity: function drawLottery(uint256 _randomSeed) returns()
func (_SingleLottery *SingleLotteryTransactor) DrawLottery(opts *bind.TransactOpts, _randomSeed *big.Int) (*types.Transaction, error) {
	return _SingleLottery.contract.Transact(opts, "drawLottery", _randomSeed)
}

// DrawLottery is a paid mutator transaction binding the contract method 0xe303a92b.
//
// Solidity: function drawLottery(uint256 _randomSeed) returns()
func (_SingleLottery *SingleLotterySession) DrawLottery(_randomSeed *big.Int) (*types.Transaction, error) {
	return _SingleLottery.Contract.DrawLottery(&_SingleLottery.TransactOpts, _randomSeed)
}

// DrawLottery is a paid mutator transaction binding the contract method 0xe303a92b.
//
// Solidity: function drawLottery(uint256 _randomSeed) returns()
func (_SingleLottery *SingleLotteryTransactorSession) DrawLottery(_randomSeed *big.Int) (*types.Transaction, error) {
	return _SingleLottery.Contract.DrawLottery(&_SingleLottery.TransactOpts, _randomSeed)
}

// EmergencyRefund is a paid mutator transaction binding the contract method 0x16bfe25c.
//
// Solidity: function emergencyRefund() returns()
func (_SingleLottery *SingleLotteryTransactor) EmergencyRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleLottery.contract.Transact(opts, "emergencyRefund")
}

// EmergencyRefund is a paid mutator transaction binding the contract method 0x16bfe25c.
//
// Solidity: function emergencyRefund() returns()
func (_SingleLottery *SingleLotterySession) EmergencyRefund() (*types.Transaction, error) {
	return _SingleLottery.Contract.EmergencyRefund(&_SingleLottery.TransactOpts)
}

// EmergencyRefund is a paid mutator transaction binding the contract method 0x16bfe25c.
//
// Solidity: function emergencyRefund() returns()
func (_SingleLottery *SingleLotteryTransactorSession) EmergencyRefund() (*types.Transaction, error) {
	return _SingleLottery.Contract.EmergencyRefund(&_SingleLottery.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleLottery *SingleLotteryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SingleLottery.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleLottery *SingleLotterySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleLottery.Contract.TransferOwnership(&_SingleLottery.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleLottery *SingleLotteryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleLottery.Contract.TransferOwnership(&_SingleLottery.TransactOpts, newOwner)
}

// SingleLotteryLotteryDrawnIterator is returned from FilterLotteryDrawn and is used to iterate over the raw logs and unpacked data for LotteryDrawn events raised by the SingleLottery contract.
type SingleLotteryLotteryDrawnIterator struct {
	Event *SingleLotteryLotteryDrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleLotteryLotteryDrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleLotteryLotteryDrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleLotteryLotteryDrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleLotteryLotteryDrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleLotteryLotteryDrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleLotteryLotteryDrawn represents a LotteryDrawn event raised by the SingleLottery contract.
type SingleLotteryLotteryDrawn struct {
	Winner      common.Address
	WinnerPrize *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLotteryDrawn is a free log retrieval operation binding the contract event 0xe270ee2f324e081c14957e0abdb70aa14e5aacd0466ac05a9a1e89982cdb5326.
//
// Solidity: event LotteryDrawn(address winner, uint256 winnerPrize)
func (_SingleLottery *SingleLotteryFilterer) FilterLotteryDrawn(opts *bind.FilterOpts) (*SingleLotteryLotteryDrawnIterator, error) {

	logs, sub, err := _SingleLottery.contract.FilterLogs(opts, "LotteryDrawn")
	if err != nil {
		return nil, err
	}
	return &SingleLotteryLotteryDrawnIterator{contract: _SingleLottery.contract, event: "LotteryDrawn", logs: logs, sub: sub}, nil
}

// WatchLotteryDrawn is a free log subscription operation binding the contract event 0xe270ee2f324e081c14957e0abdb70aa14e5aacd0466ac05a9a1e89982cdb5326.
//
// Solidity: event LotteryDrawn(address winner, uint256 winnerPrize)
func (_SingleLottery *SingleLotteryFilterer) WatchLotteryDrawn(opts *bind.WatchOpts, sink chan<- *SingleLotteryLotteryDrawn) (event.Subscription, error) {

	logs, sub, err := _SingleLottery.contract.WatchLogs(opts, "LotteryDrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleLotteryLotteryDrawn)
				if err := _SingleLottery.contract.UnpackLog(event, "LotteryDrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLotteryDrawn is a log parse operation binding the contract event 0xe270ee2f324e081c14957e0abdb70aa14e5aacd0466ac05a9a1e89982cdb5326.
//
// Solidity: event LotteryDrawn(address winner, uint256 winnerPrize)
func (_SingleLottery *SingleLotteryFilterer) ParseLotteryDrawn(log types.Log) (*SingleLotteryLotteryDrawn, error) {
	event := new(SingleLotteryLotteryDrawn)
	if err := _SingleLottery.contract.UnpackLog(event, "LotteryDrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleLotteryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SingleLottery contract.
type SingleLotteryOwnershipTransferredIterator struct {
	Event *SingleLotteryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleLotteryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleLotteryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleLotteryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleLotteryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleLotteryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleLotteryOwnershipTransferred represents a OwnershipTransferred event raised by the SingleLottery contract.
type SingleLotteryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleLottery *SingleLotteryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SingleLotteryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SingleLottery.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SingleLotteryOwnershipTransferredIterator{contract: _SingleLottery.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleLottery *SingleLotteryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SingleLotteryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SingleLottery.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleLotteryOwnershipTransferred)
				if err := _SingleLottery.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleLottery *SingleLotteryFilterer) ParseOwnershipTransferred(log types.Log) (*SingleLotteryOwnershipTransferred, error) {
	event := new(SingleLotteryOwnershipTransferred)
	if err := _SingleLottery.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleLotteryRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the SingleLottery contract.
type SingleLotteryRewardClaimedIterator struct {
	Event *SingleLotteryRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleLotteryRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleLotteryRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleLotteryRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleLotteryRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleLotteryRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleLotteryRewardClaimed represents a RewardClaimed event raised by the SingleLottery contract.
type SingleLotteryRewardClaimed struct {
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address claimer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*SingleLotteryRewardClaimedIterator, error) {

	logs, sub, err := _SingleLottery.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &SingleLotteryRewardClaimedIterator{contract: _SingleLottery.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address claimer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *SingleLotteryRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _SingleLottery.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleLotteryRewardClaimed)
				if err := _SingleLottery.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address claimer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) ParseRewardClaimed(log types.Log) (*SingleLotteryRewardClaimed, error) {
	event := new(SingleLotteryRewardClaimed)
	if err := _SingleLottery.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleLotteryTicketPurchasedIterator is returned from FilterTicketPurchased and is used to iterate over the raw logs and unpacked data for TicketPurchased events raised by the SingleLottery contract.
type SingleLotteryTicketPurchasedIterator struct {
	Event *SingleLotteryTicketPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleLotteryTicketPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleLotteryTicketPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleLotteryTicketPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleLotteryTicketPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleLotteryTicketPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleLotteryTicketPurchased represents a TicketPurchased event raised by the SingleLottery contract.
type SingleLotteryTicketPurchased struct {
	Buyer  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTicketPurchased is a free log retrieval operation binding the contract event 0x0668f5b446eb814fe35b3206f43f14bd8567ba04ddaf7a3ee56516929ab22ccb.
//
// Solidity: event TicketPurchased(address buyer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) FilterTicketPurchased(opts *bind.FilterOpts) (*SingleLotteryTicketPurchasedIterator, error) {

	logs, sub, err := _SingleLottery.contract.FilterLogs(opts, "TicketPurchased")
	if err != nil {
		return nil, err
	}
	return &SingleLotteryTicketPurchasedIterator{contract: _SingleLottery.contract, event: "TicketPurchased", logs: logs, sub: sub}, nil
}

// WatchTicketPurchased is a free log subscription operation binding the contract event 0x0668f5b446eb814fe35b3206f43f14bd8567ba04ddaf7a3ee56516929ab22ccb.
//
// Solidity: event TicketPurchased(address buyer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) WatchTicketPurchased(opts *bind.WatchOpts, sink chan<- *SingleLotteryTicketPurchased) (event.Subscription, error) {

	logs, sub, err := _SingleLottery.contract.WatchLogs(opts, "TicketPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleLotteryTicketPurchased)
				if err := _SingleLottery.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTicketPurchased is a log parse operation binding the contract event 0x0668f5b446eb814fe35b3206f43f14bd8567ba04ddaf7a3ee56516929ab22ccb.
//
// Solidity: event TicketPurchased(address buyer, uint256 amount)
func (_SingleLottery *SingleLotteryFilterer) ParseTicketPurchased(log types.Log) (*SingleLotteryTicketPurchased, error) {
	event := new(SingleLotteryTicketPurchased)
	if err := _SingleLottery.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
