package types

// Нормализовано под ER: без старых полей и алиасов.

type CommiterTxBody struct {
	Type           string `json:"type"`             // "commiter"
	ID             string `json:"id"`               // "commiter:<base64(pubkey)>"
	Name           string `json:"name"`
	CommiterPubKey string `json:"commiter_pubkey"`  // base64
}

type BeneficiaryTxBody struct {
	Type string `json:"type"` // "beneficiary"
	ID   string `json:"id"`   // "beneficiary:<uuid>"
	Name string `json:"name"`
}

type PromiseTxBody struct {
	Type            string  `json:"type"`               // "promise"
	ID              string  `json:"id"`                 // "promise:<uuid>"
	Text            string  `json:"text"`               // обяз.
	Due             int64   `json:"due"`                // unix seconds, обяз.
	BeneficiaryID   string  `json:"beneficiary_id"`     // "beneficiary:<uuid>", обяз.
	ParentPromiseID *string `json:"parent_promise_id"`  // "promise:<uuid>", опц.
}

type CommitmentTxBody struct {
	Type       string `json:"type"`        // "commitment"
	ID         string `json:"id"`          // "commitment:<uuid>"
	PromiseID  string `json:"promise_id"`  // "promise:<uuid>"
	CommiterID string `json:"commiter_id"` // "commiter:<base64(pubkey)>"
	Due        int64  `json:"due"`         // unix seconds, обяз.
}

type CompoundTx struct {
	Body struct {
		Promise    *PromiseTxBody    `json:"promise"`    // обяз.
		Commitment *CommitmentTxBody `json:"commitment"` // обяз.
	} `json:"body"`
	Signature string `json:"signature"` // base64(ed25519(body_json))
}

type SignedTx struct {
	Body      any    `json:"body"`
	Signature string `json:"signature"` // base64
}
