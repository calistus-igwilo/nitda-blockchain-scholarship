// example of proof-of-work algorithm
// Author: 70RP3D0.

const { createHash } = require('crypto');

// BigInt() not really neccessary till exponent greater than 52 eg. 2^54
// const MAX_NONCE = BigInt(Math.pow(2, 32));
const MAX_NONCE = Math.pow(2, 32);
const MAX_DIFFICULTY_LEVEL = 4;	//	Maximum leading ZEROs.

function proofOfWork(block_data, leading_zeros) {

	let block_hash = "";
	let nonce = 0;

	// Line below is commented out to allow for consistent result across multiple runs.
	block_data = block_data + Date.now();

	// for (const nonce of [...Array(MAX_NONCE).keys()]) {
	// for (const nonce of [...new Int8Array(MAX_NONCE).keys()]) {
	for (nonce = 0; nonce < MAX_NONCE; nonce++) {
		const block_header = block_data + nonce;
		block_hash = createHash('sha256').update(block_header).digest('hex');
		// console.log("[-] Hash: ", block_hash, " :", block_hash.length);

		if (block_hash.startsWith(leading_zeros)) {
			console.log("[+] Success with nonce ", nonce);
			console.log("[+] Hash is ", block_hash);
			return [block_hash, nonce];
		}
	}

	console.log(`[X] Failed after ${nonce} nonce tries ...Remove ${nonce} constraint | Update timestamp | Re-Order transactions ... then start again. `);
	return [block_hash, nonce];
}



function main() {

	let previous_block_hash = "";

	for (let difficulty_level = 1; difficulty_level <= MAX_DIFFICULTY_LEVEL; difficulty_level++) {
		console.log(`\n[-] Difficulty: ${difficulty_level} hex digits, ${difficulty_level * 8} bits`);
		console.log("[-] xTarTing search...");

		let block_content = previous_block_hash + "New Block Transactions ... ";
		const leading_zeros = "0".repeat(difficulty_level);

		// checkpoint the current time
		const start_time = Date.now();

		const [block_hash, nonce] = proofOfWork(block_content, leading_zeros);

		previous_block_hash = block_hash;

		// checkpoint how long it took to find a result
		const end_time = Date.now();

		const elapsed_time = end_time - start_time;
		console.log("[-] Elapsed Time: ", elapsed_time / 1000, " seconds");

		if (elapsed_time > 0) {
			// estimate the hashes per second
			const hash_power = ((nonce / elapsed_time));
			console.log(`[-] Hashing Power(Avg): ${hash_power} hashes per milisecond\n`);
		}
		// break;
	}
}


main();