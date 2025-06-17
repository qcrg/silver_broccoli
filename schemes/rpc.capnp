@0xd0de1bc908cbe6e6;

using Go = import "/go.capnp";

$Go.package("api");
$Go.import("api");

using Token = Data;
using BID = Int64;
using ID = Int32;
using Amount = Int64;

struct GetHistoryInfo {
  token @0 :Token;
}

struct GetHistoryResponse {
}

struct FreezeWalletInfo {
  token @0 :Token;
}

struct FreezeWalletResponse {
}

struct UnfreezeWalletInfo {
  token @0 :Token;
}

struct UnfreezeWalletResponse {
}

struct AddInfo {
  token @0 :Token;
}

struct AddResponse {
}

struct ReduceInfo {
  token @0 :Token;
}

struct ReduceResponse {
}

struct TransferInfo {
  token @0 :Token;
}

struct TransferResponse {
}

struct ReserveInfo {
  token @0 :Token;
}

struct ReserveResponse {
}

interface SilverBroccoli {
  getBalance @0 (token :Token, walletId :BID) -> (balance :Int64);
  getHistory @1 (info :GetHistoryInfo) -> (respones :GetHistoryResponse);
  createWallet @2 (token :Token, typeId :ID) -> (walletId :Int64);
  freezeWallet @3 (token :Token, walletId :BID) -> (response :Void);
  unfreezeWallet @4 (token :Token, walletId :BID) -> (response :Void);

  add @5 (token :Token, walletId :BID, amount :Amount) -> (response :AddResponse);
  reduce @6 (token :Token, walletId :BID, amount :Amount) -> (response :ReduceResponse);
  transfer @7 (token :Token, dstWalletId :BID, srcWalletId :BID, amount :Amount) -> (response :TransferResponse);
  reserve @8 (token :Token, walletId :BID, amount :Amount) -> (response :ReserveResponse);
}
