(window.webpackJsonp=window.webpackJsonp||[]).push([[115],{675:function(e,t,s){"use strict";s.r(t);var a=s(1),i=Object(a.a)({},(function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[s("h1",{attrs:{id:"concepts"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#concepts"}},[e._v("#")]),e._v(" Concepts")]),e._v(" "),s("h2",{attrs:{id:"client-state"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#client-state"}},[e._v("#")]),e._v(" Client State")]),e._v(" "),s("p",[e._v("The "),s("code",[e._v("ClientState")]),e._v(" for a solo machine light client stores the latest sequence, the frozen sequence,\nthe latest consensus state, and client flag indicating if the client should be allowed to be updated\nafter a governance proposal.")]),e._v(" "),s("p",[e._v("If the client is not frozen then the frozen sequence is 0.")]),e._v(" "),s("h2",{attrs:{id:"consensus-state"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#consensus-state"}},[e._v("#")]),e._v(" Consensus State")]),e._v(" "),s("p",[e._v("The consensus states stores the public key, diversifier, and timestamp of the solo machine light client.")]),e._v(" "),s("p",[e._v("The diversifier is used to prevent accidental misbehaviour if the same public key is used across\ndifferent chains with the same client identifier. It should be unique to the chain the light client\nis used on.")]),e._v(" "),s("h2",{attrs:{id:"public-key"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#public-key"}},[e._v("#")]),e._v(" Public Key")]),e._v(" "),s("p",[e._v("The public key can be a single public key or a multi-signature public key. The public key type used\nmust fulfill the tendermint public key interface (this will become the SDK public key interface in the\nnear future). The public key must be registered on the application codec otherwise encoding/decoding\nerrors will arise. The public key stored in the consensus state is represented as a protobuf "),s("code",[e._v("Any")]),e._v(".\nThis allows for flexibility in what other public key types can be supported in the future.")]),e._v(" "),s("h2",{attrs:{id:"counterparty-verification"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#counterparty-verification"}},[e._v("#")]),e._v(" Counterparty Verification")]),e._v(" "),s("p",[e._v("The solo machine light client can verify counterparty client state, consensus state, connection state,\nchannel state, packet commitments, packet acknowledgements, packet receipt absence,\nand the next sequence receive. At the end of each successful verification call the light\nclient sequence number will be incremented.")]),e._v(" "),s("p",[e._v("Successful verification requires the current public key to sign over the proof.")]),e._v(" "),s("h2",{attrs:{id:"proofs"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#proofs"}},[e._v("#")]),e._v(" Proofs")]),e._v(" "),s("p",[e._v("A solo machine proof should verify that the solomachine public key signed\nover some specified data. The format for generating marshaled proofs for\nthe SDK's implementation of solo machine is as follows:")]),e._v(" "),s("ol",[s("li",[e._v("Construct the data using the associated protobuf definition and marshal it.")])]),e._v(" "),s("p",[e._v("For example:")]),e._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"ZGF0YSA6PSAmYW1wO0NsaWVudFN0YXRlRGF0YXsKICBQYXRoOiAgICAgICAgW11ieXRlKHBhdGguU3RyaW5nKCkpLAogIENsaWVudFN0YXRlOiBhbnksCn0KCmRhdGFCeiwgZXJyIDo9IGNkYy5NYXJzaGFsKGRhdGEpCg=="}}),e._v(" "),s("p",[e._v("The helper functions "),s("code",[e._v("...DataBytes()")]),e._v(" in "),s("a",{attrs:{href:"../../../../modules/light-clients/06-solomachine/proof.go"}},[e._v("proof.go")]),e._v(" handle this\nfunctionality.")]),e._v(" "),s("ol",{attrs:{start:"2"}},[s("li",[e._v("Construct the "),s("code",[e._v("SignBytes")]),e._v(" and marshal it.")])]),e._v(" "),s("p",[e._v("For example:")]),e._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"c2lnbkJ5dGVzIDo9ICZhbXA7U2lnbkJ5dGVzewogIFNlcXVlbmNlOiAgICBzZXF1ZW5jZSwKICBUaW1lc3RhbXA6ICAgdGltZXN0YW1wLAogIERpdmVyc2lmaWVyOiBkaXZlcnNpZmllciwKICBEYXRhVHlwZTogICAgQ0xJRU5ULAogIERhdGE6ICAgICAgICBkYXRhQnosCn0KCnNpZ25CeiwgZXJyIDo9IGNkYy5NYXJzaGFsKHNpZ25CeXRlcykK"}}),e._v(" "),s("p",[e._v("The helper functions "),s("code",[e._v("...SignBytes()")]),e._v(" in "),s("a",{attrs:{href:"../../../../modules/light-clients/06-solomachine/proof.go"}},[e._v("proof.go")]),e._v(" handle this functionality.\nThe "),s("code",[e._v("DataType")]),e._v(" field is used to disambiguate what type of data was signed to prevent potential\nproto encoding overlap.")]),e._v(" "),s("ol",{attrs:{start:"3"}},[s("li",[e._v("Sign the sign bytes. Embed the signatures into either "),s("code",[e._v("SingleSignatureData")]),e._v(" or "),s("code",[e._v("MultiSignatureData")]),e._v(".\nConvert the "),s("code",[e._v("SignatureData")]),e._v(" to proto and marshal it.")])]),e._v(" "),s("p",[e._v("For example:")]),e._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"c2lnLCBlcnIgOj0ga2V5LlNpZ24oc2lnbkJ6KQpzaWdEYXRhIDo9ICZhbXA7c2lnbmluZy5TaW5nbGVTaWduYXR1cmVEYXRhewogIFNpZ25hdHVyZTogc2lnLAp9Cgpwcm90b1NpZ0RhdGEgOj0gc2lnbmluZy5TaWduYXR1cmVEYXRhVG9Qcm90byhzaWdEYXRhKQpieiwgZXJyIDo9IGNkYy5NYXJzaGFsKHByb3RvU2lnRGF0YSkK"}}),e._v(" "),s("ol",{attrs:{start:"4"}},[s("li",[e._v("Construct a "),s("code",[e._v("TimestampedSignatureData")]),e._v(" and marshal it. The marshaled result can be passed in\nas the proof parameter to the verification functions.")])]),e._v(" "),s("p",[e._v("For example:")]),e._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"dGltZXN0YW1wZWRTaWduYXR1cmVEYXRhIDo9ICZhbXA7c29sb21hY2hpbmUuVGltZXN0YW1wZWRTaWduYXR1cmVEYXRhewogIFNpZ25hdHVyZURhdGE6IHNpZ0RhdGEsCiAgVGltZXN0YW1wOiBzb2xvbWFjaGluZS5UaW1lLAp9Cgpwcm9vZiwgZXJyIDo9IGNkYy5NYXJzaGFsKHRpbWVzdGFtcGVkU2lnbmF0dXJlRGF0YSkK"}}),e._v(" "),s("p",[e._v("NOTE: At the end of this process, the sequence associated with the key needs to be updated.\nThe sequence must be incremented each time proof is generated.")]),e._v(" "),s("h2",{attrs:{id:"updates-by-header"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#updates-by-header"}},[e._v("#")]),e._v(" Updates By Header")]),e._v(" "),s("p",[e._v("An update by a header will only succeed if:")]),e._v(" "),s("ul",[s("li",[e._v("the header provided is parseable to solo machine header")]),e._v(" "),s("li",[e._v("the header sequence matches the current sequence")]),e._v(" "),s("li",[e._v("the header timestamp is greater than or equal to the consensus state timestamp")]),e._v(" "),s("li",[e._v("the currently registered public key generated the proof")])]),e._v(" "),s("p",[e._v("If the update is successful:")]),e._v(" "),s("ul",[s("li",[e._v("the public key is updated")]),e._v(" "),s("li",[e._v("the diversifier is updated")]),e._v(" "),s("li",[e._v("the timestamp is updated")]),e._v(" "),s("li",[e._v("the sequence is incremented by 1")]),e._v(" "),s("li",[e._v("the new consensus state is set in the client state")])]),e._v(" "),s("h2",{attrs:{id:"updates-by-proposal"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#updates-by-proposal"}},[e._v("#")]),e._v(" Updates By Proposal")]),e._v(" "),s("p",[e._v("An update by a governance proposal will only succeed if:")]),e._v(" "),s("ul",[s("li",[e._v("the substitute provided is parseable to solo machine client state")]),e._v(" "),s("li",[e._v("the new consensus state public key does not equal the current consensus state public key")])]),e._v(" "),s("p",[e._v("If the update is successful:")]),e._v(" "),s("ul",[s("li",[e._v("the subject client state is updated to the substitute client state")]),e._v(" "),s("li",[e._v("the subject consensus state is updated to the substitute consensus state")]),e._v(" "),s("li",[e._v("the client is unfrozen (if it was previously frozen)")])]),e._v(" "),s("p",[e._v("NOTE: Previously, "),s("code",[e._v("AllowUpdateAfterProposal")]),e._v(" was used to signal the update/recovery options for the solo machine client.  However, this has now been deprecated because a code migration can overwrite the client and consensus states regardless of the value of this parameter. If governance would vote to overwrite a client or consensus state, it is likely that governance would also be willing to perform a code migration to do the same.")]),e._v(" "),s("h2",{attrs:{id:"misbehaviour"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#misbehaviour"}},[e._v("#")]),e._v(" Misbehaviour")]),e._v(" "),s("p",[e._v("Misbehaviour handling will only succeed if:")]),e._v(" "),s("ul",[s("li",[e._v("the misbehaviour provided is parseable to solo machine misbehaviour")]),e._v(" "),s("li",[e._v("the client is not already frozen")]),e._v(" "),s("li",[e._v("the current public key signed over two unique data messages at the same sequence and diversifier.")])]),e._v(" "),s("p",[e._v("If the misbehaviour is successfully processed:")]),e._v(" "),s("ul",[s("li",[e._v("the client is frozen by setting the frozen sequence to the misbehaviour sequence")])]),e._v(" "),s("p",[e._v("NOTE: Misbehaviour processing is data processing order dependent. A misbehaving solo machine\ncould update to a new public key to prevent being frozen before misbehaviour is submitted.")]),e._v(" "),s("h2",{attrs:{id:"upgrades"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#upgrades"}},[e._v("#")]),e._v(" Upgrades")]),e._v(" "),s("p",[e._v("Upgrades to solo machine light clients are not supported since an entirely different type of\npublic key can be set using normal client updates.")])],1)}),[],!1,null,null,null);t.default=i.exports}}]);