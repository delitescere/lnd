package sorceror

import (
	"github.com/boltdb/bolt"
	"github.com/btcsuite/btcd/wire"
)

const defaultpath = "."

type SorceStore struct {
	Path    string
	SorceDB *bolt.DB
}

// 2 structs that the sorceror gets from clients: Descriptors and Msgs

// SorceDescriptor is the initial / static description of a SorceChan
type SorceDescriptor struct {
	DestPKHScript [20]byte // PKH to grab to; main unique identifier.

	Delay uint16 // timeout in blocks
	Fee   int64  // fee to use for grab tx. could make variable but annoying...

	HAKDBasePoint [33]byte // client's HAKD key base point
	TimeBasePoint [33]byte // potential attacker's timeout basepoint

	// elk 0 here?  Because they won't send a sig for elk0...
	ElkZero wire.ShaHash
}

// the message describing the next state, sent from the client to the sorceror
type SorceMsg struct {
	DestPKHScript [20]byte     // identifier for channel; could be optimized
	Txid          wire.ShaHash // txid of close tx
	Elk           wire.ShaHash // elkrem for this state index
	Sig           [64]byte     // sig for the grab tx
}

// 2 structs used in the DB: IdxSigs and ChanData

// IdxSig is what we save in the DB for each txid
type IdxSig struct {
	PKHIdx   uint32
	StateIdx uint64
	Sig      [64]byte
}

// ChanStatic is data that doesn't change
type ChanStatic struct {
	Delay uint16 // timeout in blocks
	Fee   int64  // fee to use for grab tx. could make variable but annoying...

	HAKDBasePoint [33]byte // client's HAKD key base point
	TimeBasePoint [33]byte // potential attacker's timeout basepoint

	PeerIdx uint32 // can save the user you're watching this for.  Optional
}

// IngestDescriptor takes a channel descriptor and sets up the DB to
// receive state update messages
func (s *SorceStore) IngestDescriptor(d SorceDescriptor) error {
	return nil
}

// IngestMsg takes a state update message and saves it in the DB.
// A bunch of errors are possible.
func (s *SorceStore) IngestMsg(m SorceMsg) error {
	return nil
}

// Ingest the next state.  Will error half the time if the elkrem's invalid.
// Never errors on invalid sig.
//func (sc *SorceStore) Ingest(sm SorceMsg) error {
//	if sc == nil {
//		return fmt.Errorf("Ingest: nil SorcedChan")
//	}
//	// first ingest the elkrem
//	err := sc.Elk.AddNext(&sm.Elk)
//	if err != nil {
//		return err
//	}
//	// serialize elkrem
//	//	elkBytes, err := sc.Elk.ToBytes()
//	//	if err != nil {
//	//		return err
//	//	}
//	// should mv elk to oldelk here?  For faster recovery if write fails?
//	// not really critical though as this is a backup anyway and can re-sync it
//	// from the client

//	return nil
//}

// Grab produces the grab tx, if possible.
func (sc *SorceStore) Grab(cTx *wire.MsgTx) (*wire.MsgTx, error) {

	// sanity chex
	//	if sc == nil {
	//		return nil, fmt.Errorf("Grab: nil SorcedChan")
	//	}
	//	if cTx == nil || len(cTx.TxOut) == 0 {
	//		return nil, fmt.Errorf("Grab: nil close tx")
	//	}
	//	// determine state index from close tx
	//	stateIdx := uspv.GetStateIdxFromTx(cTx, sc.XorIdx)
	//	if stateIdx == 0 {
	//		// no valid state index, likely a cooperative close
	//		return nil, fmt.Errorf("Grab: close tx has 0 state index")
	//	}
	//	// check if we have sufficient elkrem
	//	if stateIdx >= sc.Elk.UpTo() {
	//		return nil, fmt.Errorf("Grab: state idx %d but elk up to %d",
	//			stateIdx, sc.Elk.UpTo())
	//	}
	//	// check if we have sufficient sig.  This is redundant because elks & sigs
	//	// should always be in sync.
	//	if stateIdx > uint64(len(sc.Sigs)) {
	//		return nil, fmt.Errorf("Grab: state idx %d but %d sigs",
	//			stateIdx, len(sc.Sigs))
	//	}
	//	PubArr := sc.BasePoint
	//	elk, err := sc.Elk.AtIndex(stateIdx)
	//	if err != nil {
	//		return nil, err
	//	}
	//	err = uspv.PubKeyArrAddBytes(&PubArr, elk.Bytes())
	//	if err != nil {
	//		return nil, err
	//	}

	//	// figure out amount to grab
	//	// for now, assumes 2 outputs.  Later, look for the largest wsh output
	//	if len(cTx.TxOut[0].PkScript) == 34 {
	//		shIdx = 0
	//	} else {
	//		shIdx = 1
	//	}

	//	// calculate script for p2wsh
	//	preScript, _ := uspv.CommitScript2(PubArr, sc.OtherRefdundPub, sc.Delay)

	//	// annoying 2-step outpoint calc
	//	closeTxid := cTx.TxSha()
	//	grabOP := wire.NewOutPoint(&closeTxid, 0)
	//	// make the txin
	//	grabTxIn := wire.NewTxIn(grabOP, nil, make([][]byte, 2))
	//	// sig, then script
	//	grabTxIn.Witness[0] = sc.Sigs[stateIdx]
	//	grabTxIn.Witness[1] = preScript

	//	// make a txout
	//	grabTxOut := wire.NewTxOut(10000, sc.DestPKHScript[:])

	//	// make the tx and add the txin and txout
	//	grabTx := wire.NewMsgTx()
	//	grabTx.AddTxIn(grabTxIn)
	//	grabTx.AddTxOut(grabTxOut)

	//	return grabTx, nil
	return nil, nil
}
