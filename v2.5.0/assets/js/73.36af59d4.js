(window.webpackJsonp=window.webpackJsonp||[]).push([[73],{585:function(e,t,a){"use strict";a.r(t);var n=a(0),i=Object(n.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"ibc-client-developer-guide-to-upgrades"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#ibc-client-developer-guide-to-upgrades"}},[e._v("#")]),e._v(" IBC Client Developer Guide to Upgrades")]),e._v(" "),a("p",{attrs:{synopsis:""}},[e._v("Learn how to implement upgrade functionality for your custom IBC client.")]),e._v(" "),a("p",[e._v("As mentioned in the "),a("RouterLink",{attrs:{to:"/ibc/upgrades/"}},[e._v("README")]),e._v(", it is vital that high-value IBC clients can upgrade along with their underlying chains to avoid disruption to the IBC ecosystem. Thus, IBC client developers will want to implement upgrade functionality to enable clients to maintain connections and channels even across chain upgrades.")],1),e._v(" "),a("p",[e._v("The IBC protocol allows client implementations to provide a path to upgrading clients given the upgraded client state, upgraded consensus state and proofs for each.")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gVXBncmFkZSBmdW5jdGlvbnMKLy8gTk9URTogcHJvb2YgaGVpZ2h0cyBhcmUgbm90IGluY2x1ZGVkIGFzIHVwZ3JhZGUgdG8gYSBuZXcgcmV2aXNpb24gaXMgZXhwZWN0ZWQgdG8gcGFzcyBvbmx5IG9uIHRoZSBsYXN0Ci8vIGhlaWdodCBjb21taXR0ZWQgYnkgdGhlIGN1cnJlbnQgcmV2aXNpb24uIENsaWVudHMgYXJlIHJlc3BvbnNpYmxlIGZvciBlbnN1cmluZyB0aGF0IHRoZSBwbGFubmVkIGxhc3QKLy8gaGVpZ2h0IG9mIHRoZSBjdXJyZW50IHJldmlzaW9uIGlzIHNvbWVob3cgZW5jb2RlZCBpbiB0aGUgcHJvb2YgdmVyaWZpY2F0aW9uIHByb2Nlc3MuCi8vIFRoaXMgaXMgdG8gZW5zdXJlIHRoYXQgbm8gcHJlbWF0dXJlIHVwZ3JhZGVzIG9jY3VyLCBzaW5jZSB1cGdyYWRlIHBsYW5zIGNvbW1pdHRlZCB0byBieSB0aGUgY291bnRlcnBhcnR5Ci8vIG1heSBiZSBjYW5jZWxsZWQgb3IgbW9kaWZpZWQgYmVmb3JlIHRoZSBsYXN0IHBsYW5uZWQgaGVpZ2h0LgpWZXJpZnlVcGdyYWRlQW5kVXBkYXRlU3RhdGUoCiAgICBjdHggc2RrLkNvbnRleHQsCiAgICBjZGMgY29kZWMuQmluYXJ5Q29kZWMsCiAgICBzdG9yZSBzZGsuS1ZTdG9yZSwKICAgIG5ld0NsaWVudCBDbGllbnRTdGF0ZSwKICAgIG5ld0NvbnNTdGF0ZSBDb25zZW5zdXNTdGF0ZSwKICAgIHByb29mVXBncmFkZUNsaWVudCwKICAgIHByb29mVXBncmFkZUNvbnNTdGF0ZSBbXWJ5dGUsCikgKHVwZ3JhZGVkQ2xpZW50IENsaWVudFN0YXRlLCB1cGdyYWRlZENvbnNlbnN1cyBDb25zZW5zdXNTdGF0ZSwgZXJyIGVycm9yKQo="}}),e._v(" "),a("p",[e._v("Note that the clients should have prior knowledge of the merkle path that the upgraded client and upgraded consensus states will use. The height at which the upgrade has occurred should also be encoded in the proof. The Tendermint client implementation accomplishes this by including an "),a("code",[e._v("UpgradePath")]),e._v(" in the ClientState itself, which is used along with the upgrade height to construct the merkle path under which the client state and consensus state are committed.")]),e._v(" "),a("p",[e._v("Developers must ensure that the "),a("code",[e._v("UpgradeClientMsg")]),e._v(" does not pass until the last height of the old chain has been committed, and after the chain upgrades, the "),a("code",[e._v("UpgradeClientMsg")]),e._v(" should pass once and only once on all counterparty clients.")]),e._v(" "),a("p",[e._v("Developers must ensure that the new client adopts all of the new Client parameters that must be uniform across every valid light client of a chain (chain-chosen parameters), while maintaining the Client parameters that are customizable by each individual client (client-chosen parameters) from the previous version of the client.")]),e._v(" "),a("p",[e._v("Upgrades must adhere to the IBC Security Model. IBC does not rely on the assumption of honest relayers for correctness. Thus users should not have to rely on relayers to maintain client correctness and security (though honest relayers must exist to maintain relayer liveness). While relayers may choose any set of client parameters while creating a new "),a("code",[e._v("ClientState")]),e._v(", this still holds under the security model since users can always choose a relayer-created client that suits their security and correctness needs or create a Client with their desired parameters if no such client exists.")]),e._v(" "),a("p",[e._v("However, when upgrading an existing client, one must keep in mind that there are already many users who depend on this client's particular parameters. We cannot give the upgrading relayer free choice over these parameters once they have already been chosen. This would violate the security model since users who rely on the client would have to rely on the upgrading relayer to maintain the same level of security. Thus, developers must make sure that their upgrade mechanism allows clients to upgrade the chain-specified parameters whenever a chain upgrade changes these parameters (examples in the Tendermint client include "),a("code",[e._v("UnbondingPeriod")]),e._v(", "),a("code",[e._v("TrustingPeriod")]),e._v(", "),a("code",[e._v("ChainID")]),e._v(", "),a("code",[e._v("UpgradePath")]),e._v(", etc.), while ensuring that the relayer submitting the "),a("code",[e._v("UpgradeClientMsg")]),e._v(" cannot alter the client-chosen parameters that the users are relying upon (examples in Tendermint client include "),a("code",[e._v("TrustLevel")]),e._v(", "),a("code",[e._v("MaxClockDrift")]),e._v(", etc).")]),e._v(" "),a("p",[e._v("Developers should maintain the distinction between Client parameters that are uniform across every valid light client of a chain (chain-chosen parameters), and Client parameters that are customizable by each individual client (client-chosen parameters); since this distinction is necessary to implement the "),a("code",[e._v("ZeroCustomFields")]),e._v(" method in the "),a("code",[e._v("ClientState")]),e._v(" interface:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gVXRpbGl0eSBmdW5jdGlvbiB0aGF0IHplcm9lcyBvdXQgYW55IGNsaWVudCBjdXN0b21pemFibGUgZmllbGRzIGluIGNsaWVudCBzdGF0ZQovLyBMZWRnZXIgZW5mb3JjZWQgZmllbGRzIGFyZSBtYWludGFpbmVkIHdoaWxlIGFsbCBjdXN0b20gZmllbGRzIGFyZSB6ZXJvIHZhbHVlcwovLyBVc2VkIHRvIHZlcmlmeSB1cGdyYWRlcwpaZXJvQ3VzdG9tRmllbGRzKCkgQ2xpZW50U3RhdGUK"}}),e._v(" "),a("p",[e._v("Counterparty clients can upgrade securely by using all of the chain-chosen parameters from the chain-committed "),a("code",[e._v("UpgradedClient")]),e._v(" and preserving all of the old client-chosen parameters. This enables chains to securely upgrade without relying on an honest relayer, however it can in some cases lead to an invalid final "),a("code",[e._v("ClientState")]),e._v(" if the new chain-chosen parameters clash with the old client-chosen parameter. This can happen in the Tendermint client case if the upgrading chain lowers the "),a("code",[e._v("UnbondingPeriod")]),e._v(" (chain-chosen) to a duration below that of a counterparty client's "),a("code",[e._v("TrustingPeriod")]),e._v(" (client-chosen). Such cases should be clearly documented by developers, so that chains know which upgrades should be avoided to prevent this problem. The final upgraded client should also be validated in "),a("code",[e._v("VerifyUpgradeAndUpdateState")]),e._v(" before returning to ensure that the client does not upgrade to an invalid "),a("code",[e._v("ClientState")]),e._v(".")])],1)}),[],!1,null,null,null);t.default=i.exports}}]);