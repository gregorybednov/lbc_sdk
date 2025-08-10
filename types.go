package types

// Types module
// Can be used for an import
// TODO someday it could be migrated onto generated types


type CommiterTxBody struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	CommiterPubKey string `json:"commiter_pubkey"`
}

type BeneficiaryTxBody struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PromiseTxBody struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Timestamp   int64  `json:"timestamp,omitempty"` // ← чтобы понимать клиент
	Title       string `json:"title,omitempty"`     // ← опционально, если когда-нибудь пригодится
	Deadline    string `json:"deadline,omitempty"`
}

type CommitmentTxBody struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	PromiseID   string `json:"promise_id"`
	CommiterID  string `json:"commiter_id"`
	CommiterSig string `json:"commiter_sig,omitempty"`
}

type CompoundTx struct {
	Body struct {
		Promise    *PromiseTxBody    `json:"promise"`
		Commitment *CommitmentTxBody `json:"commitment"`
	} `json:"body"`
	Signature string `json:"signature"`
}

type SignedTx struct {
	Body      any    `json:"body"` // TODO someday it should become less abstract
	Signature string `json:"signature"`
}


