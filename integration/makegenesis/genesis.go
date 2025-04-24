package makegenesis

import (
	"math/big"
	"time"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/cryptod"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/driver"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/driverauth"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/evmwriter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/gpos"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/netinit"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/sfc"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesisstore"
)

var (
	//FakeGenesisTime = inter.Timestamp(1608600000 * time.Second)
	FakeGenesisTime = inter.Timestamp(1642595720 * time.Second)
)

// FakeKey gets n-th fake private key.
/* func FakeKey(n int) *ecdsa.PrivateKey {

	fmt.Println("FakeKey(n int)", "testing")

	reader := rand.New(rand.NewSource(int64(n)))

	key, err := ecdsa.GenerateKey(crypto.S256(), reader)
	if err != nil {
		panic(err)
	}

	return key
} */

// FakeKey gets n-th fake private key.
func FakeKey(n int) *cryptod.PrivateKey {

	key, err := cryptod.GenerateMLDsa87Key()
	if err != nil {
		panic(err)
	}

	return key
}

/* func FakeKey(n int) *cryptod.PrivateKey {

	pk, something := cryptod.HexToMLDsa87("2ff5fb087d3d25e1979437f9a39628c0b894b77041d3660da831ba735f5a0e8480a1dc2052ac9374f28e34681f715b2ecac8e09f050d00b4d38e47edc16f6d6265676c8e009c26957f86280e225cc56cfca60efa18e106ca31c4a83d022370ac135dfd902619298f19a1997a771cc0318897916f7a08c769084782cb617449e952465243c00009888c0ac78820020e0484611ab4318ca06cdb10080c8410092580501884939025a118628ac4905a902d82c84198b20919106de048221c98295c100acb12526142321b144883884d89a460a09000d0001188b270d3b84cc21044c0289120a12484006e49b2850a168c4138885090441433466332885498240c3864db364e10146a589605c9248614156284406ed0928d9b28824088481b208c14225124936ca1462ada4885891420caa881d2c67003b12418c18400188012468d598028c90884123201cc362100a8050b04641c43061a4364d024068c2060c81071638281dc20012206421b263199209109492ea020228ac268594429103642140468a084419b866c0bc444cbb645840686233470144970d8384d1381848b885084020c4b922061867103996ca1b4485c206008396054066808c5698248010011890c2600c19850d100010c126652841090422663202519b09110b07108097100b231a428052218121b31040bc94ce2066ed2b6481913244a1826c82081d990640a2588c4a208213991604605d116664b020c18b74d1ab5251ab2600b490a49048c8102021b8369091442e4045240320d04a5880c0311d82242d0c06492008e18118098900d54462891406518152e8c26665a0292c2284c03152e6240711382604c48061aa720d1a84c9b8024c4046ec2102e092601c206684b300e02b590e288801c8930d38244db200a494026c22428c2029193328e00919064444019376a61964890966482440918a270190442c980716208261122418810486120710a18849c8410924692104728cb2806a2060a81307181460e1228690a350cd4c251612466d8c48d6394841cc26c64a020e0140421a5440837890b134cd4360c59324c10c82d1199655a86118c328d431272134786203544c0206c81201181164823372a0b19025884651bc86450186920398504a4101a9385ca96481b9929203628d4061212194cdcb02d921481a2c681c8a860cab6299b082109a224a4466858c6698004296030714c968882103091b850a106526048891c188e10254c133771090181c34461cbc08412806d214085c93431c9c08c22c4418a082e19265209b44cd2400422b56489b62c11483249c089098788438601c4922cc0129143184541383121148952822d58944cdc2046e39681190272941092432631d1446464a2108c000581b250d282455c2606e4a44dc9162960386a538270d2246a12024962b0510221011b9684ca388d1c8481a2149258326e92461158a28d0a88915c062a22876dd834054a083284264141a8109c3846cc462ee12671a1847014a94090b4411b8481cc14521329415426298a064edb92241208880c05098c444dcbc0840a032dc320411034728a16658380911c0772d0a02404276809150e1ca18142a63021834400114d8834320b980993a62842440444080a8b244209192918c09120c868224624ca0609e01680e114520b858048c269813801c4c02c8a108609940092344443b84990108289046c831666dba0411c258dd4326401a7204a805122c1690282844c94091281495134680917414a001188a404d11608134065ccc67160168c1482481a380d233462e4c62518860512c1688cc42d5c926d84c88c23b848d1928412458153826824b490438870483021934862da482011278094c68ce4922102860cc4c06820474c0c91409a46662415204902061a00485aa6059b0210db3848dc02080202095a3882e124292433020326229b8444c3223294000e82a4688020010017218a088c48383298c2681846858830499b1672ca88418b466512496ed8922092c0081c038ea3b6244ac41191b42423326d8812891995104a48609390008ba6100bc94400332c21902559340461220d524240dc802100122ecb942461884592060420131202a71113904808380421328dca089261a6682441265b1625000240cb982059000193982c23b46803a92c4b3889c4c86d929825003082041460e00200a3702cbe4681ccc0e7aacda0ef79c31dbc88cf9d3cc041783150e2959b9cead6147180240afeb5d0208331d72be162c4825fe845ad264eae6cf54d7eeddda9a6f3901b034900ae2b560b1ad140509a017e79175de67a4df54bbbfccbeab47a6f317740530fb888f4d987d004a65d88145535b1e984bc58498763ab2ba7b702c104820212baafb3d420c8493d71bd0bb8fa86e85b13a172e30ff185d3a494b798f30b661e9ac0553d0748629b003032dfaf937043576b4ae141fe9de2b5a4c36023937c2f23868530fcb9a0dd1bc36d37dcfff3adae65118fe3e335cd6d192e0c4b895e456b2a050cc52f5009b7f2400f655cab1c72edc742850a3e2b5a18f61dceab7609f48d40edf328ba6485312acb36ddb9f76e0e500148fd88a4b5465388c6a45674c121c826ff3d54dd4d950e6424b881b40aa98e3eab26a80051a1ef07b9824c8e3b1c65bed1b3f86fab015ba7032d2fa185b4c22274e7c253b145a9f609f3240d4bac2e0d29d1981510845f1487cc33179b440141894e4cce76fd2d595c1f679ed8848cee72857de1ce320992ce4151a30fb744c226d6fdc14679bbc671ceadaee1a1375caff9f499a746ec7813649a81491648062140e36e40745cf1120a04fc1ef7c4f0c1ed36d14ac9f15f0e9dadf712c8c68003589c2e64863600fd1267fc64558008b0c240fc7a03adea80c08a28cd6a29c987914ef69697fc3efbef24383705e4377f99f25196d2c2de6d50239232b9e7253f85684f822f591f69fee4803251d0a75e582ddd57dc6e8eeec4787620c66fa54a9635bec18847c5877989dd4e4ce1a2830dff312e5d6c2c5b596afe921e4df73319ca736668360f063c7334609a215491279823ff544badc17dca82096a23eaa6baa83aba9fe741d021f2531644a5fd8ca643ff087511764f1d917b4490103fc03764fcef1f99588d9efb3cfb6b92d35b3dfe3d00b7bce71de6e93e74cc278ee7c66499f1bbeedbd8f344cd8badf510b4ca31f60ffe491eaac7448aa978c857b7fa63a446b6af3b63215c41deb3c69b30f8acf389ef3959a5d566a91b9245e45b4b519cff3b07088d77ab2dc27ce3c513a8077966b09c0ea9b08f10e338dccb2fe8f9a7b4c83f4ad3dcffa24439d7e4c2ada35b26b12b6c7a3304d203451d165222ee895abf34858611fe2cdc88f3c859a178faa5f5604087cc4a1660df515bd8dca5b8e6db8318b5515d34bc003181c39c3185f234acb0889b7094c1e457b76619e16a64a9c11016383845afc3b94c7cdca0d9ed6a866bd3ae1c410fd06da8ded941ee4b832177781c269dc73b06d340599a5398d73727128c66f4796d30091540f1c7802723c8a8c4a4a49be3b9f1b0cb8d2d953aa4db85380f0a1f1081162a10fa55442099599f3ff0da6f65f1205cdd277c3f73afdd356b4e499eb87184791ad8f6d78a49409b2893bb66125aa424846461ba79502d7e09e2b5da238480877128406440d252876e7fdadae80e8d72a350242bd28c311733f2c98342f1ee099db576ba54d10e9fc53822edcb5a5895b83f8cac3e12bc6d3e0d0e22d95f7c143a7ba31fd899e6aa8899e1169af7960be2b4d449a56a6b2d6d2e536f855de23fcaced65dbb8d6ad07424755492143276233ebf67119e300fc63a70258050a18b9c4c85ae3757450737474a257effeb40f23e96ca67628ddb44ca967a470842ef231fe023db39d6856c9ae6c126f995e73b7be3368bf13accf2df7ddfab5a2e021d21d9cf6cca4d86fa7aa363aee24230337039137e77b028c879c11e1f987b012ff6894f71c4c0a984bc8a3f34acb968111c1f9c29a9bd53f5a81747a0bc10992cc7c94c66b2343ae825fa29057e3490f976c9838d80cb0f9a723f1e81982a63f53f0d94592f3f4412f276e10decb15e09b0dae1341717a7a5752c334efde93975f3fd2bccbc03beb9ead1a4551513352375988c59eddd139b6438aaaa15ebdcc9062ce1a196b231f5ccbfe2548fbae69155569f0fb16b3fd8527dc19bcf7916c59cd0b82c4727abbb32ab437586ee2c72c4b51644a563b7f1159792a45138f78a721ff6fdef138734f916ed459b5437b2e8553652bddbfd7f8c8b38cf82837abcea0719dde858dd0cec2ebc9b184d9809f9a43d8ff52377d20c4fb6defe4fc6618351aaa082a9b743ac25677682785b0c67d07e5ceb53059b7f23004cdbfa70e1057042d19767a93819cd06358dd6429b7c137c1a97f99420ffd5721391ccce9d5a2553b7620f977eab929ff53d55f1571e609be8ca26154b3058fcea8580ac03a386bd9374e524f192dbfdc005a0e1fd2ac11315657b7875d9b27535593830f097399a39a919bb448b84227a49a297c5f6ebc61d97527ec570aa9ba45d660c67f8cd75045774119477c1bfb38a82c3c4700d197d14096b5faefa2b66870c515c8b5a5dc277b89e7bb03ed85171c211c72057413b5f97a03ab6712e38c37c85e1e9ef3ddc58bdbf87b66e07456fb99a972491940d78a18ecc020797019712b8cee81773c94d2131d1586ca502672e080e243ec635f40f457094c80cd12b5016f6d07959a212c94e2b8ceaf6bdbcc9c4156bfe81081a1c1fdf5cb8010a992c51f54032fa1259464c5fc7980de98e6427643b2bf7e163d7b038adf47f054be034769f2418ff0fc096980e49e980331a14593b9098e219a036a8b8ead6d509be0b409aac4966bd731b0d42365577e53b643c7f1f8d4204e31dd07ed2ec32fa303b36001e56fffc5316fc7231a433e08d69ba96a3cdadd6f1f33853661ad55eaf5fda317520e748b6ce1cb0735c850d1b5aa5c1f1a5d4caa0675d177304a1d5d75ef526596c6a319efb810d4fc05bd3c632a42c7231ed4235887a2611df4be85b4918031823966d34e0467c8597b105771cf30eae07f6bab87959f1c80a7c6bf058947851da806f8d69f1d59850875264ed28df43ffb79c9135ed3e5b5b46ec9cac2c896f3c0b588b8c71ff384e16005612007abd41f8dff2723f6964dc155fb08812223f30fd9704eaf2cb8401c9939b43903ea5ae3cfa72b11cd6a61c498de2db31d1320d64b481a3928600a2d1997b81bc3e3dad430077693e3bf56308446f192e74b4ff5269cc72ad286e670a4d77adfa3fd61796cadbc57410b878ec616c788accf714e1a7fcc3d922225adf067a978e3d5fee44d1d136cc476e8fee087b280c4897e266d1c14d576c22eccd19842dd6d2b9f5abd5fbdca551d52e6afd475227b4d059618c9204c9193c242d3278c2dd0758f9d2901b1c9297846fc67f7b279d2723d11ddda7b5149231e4ef3ee81b15d85426b01e0ae9fbf37143e03543cbaee3e8b728efd6da9298afe9a65af82274ffb2e44b97ef410947f46d99eda7bfc585eb6de823bfae2ec867d384b2bde7a1f229a3cbdd123f0888610792ed5443ff2af9c379736ca3357dc69c88b9c27b0e04c44d414c905f9eed051d2adfcc6cccdc1265cc536a2da6f8c60bf0f0e3f89164e7a1a54e57a5134434e20aad4ef3e8c2f4aaef2d4118110b821a87ba23256e74cf55f53e5ec1f573a6e8be7ece545aa9684239e5c3944db02bb85db1122bbc225c63c87ee0185e5d6c3c13c0dd279a640d748de82160336e67b84c14f584295d821735d614581b594f1c6a92d148e48fbb2e81f548ef9a007b4f2ab8ebe7607752f0b2adedc3933f477bd4e7391b399327927cb0b5a6c8886327e4763ad7bba1ad883ae9f7fe2dba045a74c32d6c3b218804fb7fac53534cf5d1d0c7bbc56ab1d27f911061af6dd8f164b08217b5258ecbbdb4d8da75a59ef2425b604f8e2ba56e56fb231e4a936e81499a6222324112c33207a58892fb6af87d80dddeec2455dab328ab33d38583ac14e135c9546d5b72ce57084a2ff9526e331dfdd4e557811be973f86a184508133859ad894ca9bd35811faf3d563365b5d585b433698408efe9966e24139b9fb01d7b0821bbc8b872c99581fe4a77e92cafe05ff88f773e3d6e1333d047250dc1196272f0fe63ed77fd2c248b3b81be0133f2be88c9b2e1500596110dfd47f94ae25f0a25d76dcbbfe3e6eaa382fae4b11f277adcef3c6e2f0dfdb2eaac41e4deb875bc452a2326227a3e27689f5bce54dd3ff80af1c4598e5468ccb70fd0ee1d078a733085ea9e847e84549ab521027b386fa0ae75c16abeecf968361cb9bd541c2c5fe9f25d5ab584db489ea8fddc56fc8a7ce22e83f095c56c010ec02e8d50c0fa27215cd57efb55eca961cbbb75efd50598bdbaca408c6bba2d99f348eedcf17ca3bc367684fbcb083864779935af5d66f1564831a7e38f13c94500ad4cb840916eb36d38b93a4371a2beed5db0ee31e63e51a048cf65a55f94a6887b98a7aaa7dc67bd1d002839319c7c68031697574efd967a9b381179477c100ce9f2a1fcc294016d72fcd5668c8f7af0566605af354bd210e0b1385d3c8d646269d5f82f36f78d796bafa09669fd43bb53a334ae1984edb7999ce599d16fdc86f32f1907aaac377865341dfca6051f03cecdf51e3830df562bac26f76326548167f3eb568df16e8c9b2057a1f97a87c977f8a15f6c4ba4cf38430a640b3bb155d449702cf50ef99e56f633f824a2202df7ee9c5c052ca47af3b78f9634b992739e6bd65567d31de1847a1deae9776c95e3e543ebeb6cd1dccf3063d346e087a14b2a14ed538c98288fff0a2c1476e4")
	//fmt.Println("Generated key", pk)
	if something != nil {
		fmt.Println("Generated something", something.Error())
	}
	return pk
} */

func FakeGenesisStore(num int, balance, stake *big.Int) *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(ncogearthchain.FakeNetRules())

	validators := GetFakeValidators(num)

	totalSupply := new(big.Int)
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          FakeGenesisTime,
		PrevEpochTime: FakeGenesisTime - inter.Timestamp(time.Hour),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        FakeGenesisTime - inter.Timestamp(time.Minute),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}

/* func GetFakeValidators(num int) gpos.Validators {
	validators := make(gpos.Validators, 0, num)

	for i := 1; i <= num; i++ {
		key := FakeKey(i)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		pubkeyraw := crypto.FromECDSAPub(&key.PublicKey)
		validatorID := idx.ValidatorID(i)
		validators = append(validators, gpos.Validator{
			ID:      validatorID,
			Address: addr,
			PubKey: validatorpk.PubKey{
				Raw:  pubkeyraw,
				Type: validatorpk.Types.Secp256k1,
			},
			CreationTime:     FakeGenesisTime,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
} */

func GetFakeValidators(num int) gpos.Validators {
	validators := make(gpos.Validators, 0, num)

	for i := 1; i <= num; i++ {
		key := FakeKey(i)
		addr := cryptod.PubkeyToAddress(*key.Public().(*cryptod.PublicKey))
		pubkeyraw := cryptod.FromMLDsa87Pub(key.Public().(*cryptod.PublicKey))
		validatorID := idx.ValidatorID(i)
		validators = append(validators, gpos.Validator{
			ID:      validatorID,
			Address: addr,
			PubKey: validatorpk.PubKey{
				Raw:  pubkeyraw,
				Type: validatorpk.Types.MLDsa87,
			},
			CreationTime:     FakeGenesisTime,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
}
