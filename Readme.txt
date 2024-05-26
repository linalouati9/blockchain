utils.go file :

function IntToHex : 
This function utilizes the "binary.BigEndian.PutUint64" function 
from the encoding/binary package to perform the conversion of 
the int64 into a byte array ([]byte).


function StartsWithXZeros :
The StartsWithXZeros algorithm checks whether the beginning of a hash,
represented as a byte array, starts with a specified number of leading
zero bits. It first examines complete bytes, ensuring they are zero, 
and then verifies any remaining bits if the specified number is not 
divisible by 8. If these conditions are met, the function returns true,
otherwise, it returns false.

function EqualSlices :
The EqualSlices function compares two byte slices, first checking if 
they have the same length. If the lengths differ, it returns false. 
Otherwise, it iterates through each byte, returning false at the first
difference found; otherwise, it returns true at the end of the 
comparison.

function EqualMaps :
The EqualMaps function follows a similar logic to EqualSlices, for more
details check comments.

function EqualTransactions:
The EqualTransactions function checks if two transactions are equal 
by comparing their hash slices using the EqualSlices function.

function EqualBlocks:
The 'EqualBlocks' function checks if two blocks are equal by comparing
their hash values using the `EqualSlices` utility function.

function Serialize :
The Serialize function converts a slice of byte slices into a single 
byte slice by appending the elements of each slice.

***********************************************************************
transaction.go file :

function NewTransaction :
The NewTransaction function create a new Transaction instance by 
initializing its fields with the provided parameters. It takes a hash,
a slice of transaction inputs (txIns), and a slice of transaction 
outputs (txOuts), then returns a pointer to the newly created 
transaction instance.

function computeHash : 
The idea behind the function is to use the Gob binary encoding to 
convert the transaction into a binary representation. Subsequently, 
it applies the SHA-256 hashing algorithm to generate the final hash 
of the transaction.

function NewCoinbaseTX :
the function encapsulates the process of creating a coinbase transaction,
taking care of setting default data, creating appropriate inputs and 
outputs, generating the transaction, and computing its hash.
***********************************************************************
persistency.go file :

function bcFileExists :
The 'bcFileExists' function check if a file exists in the file system.
It takes a file path as a parameter, uses the os.Stat function to 
obtain information about the file, and then checks whether the error 
returned indicates that the file does not exist (os.IsNotExist(err)).
The function returns true if the file exists and false otherwise.

function LoadBlockchain :
The 'LoadBlockchain' function checks if a blockchain file exists,
opens and reads its content, decodes the JSON data into a Blockchain
struct, and returns the loaded blockchain or an error.

function SaveBlockchain :
The 'SaveBlockchain' function saves a blockchain to a file by opening
the file for writing, encoding the blockchain data to JSON, and 
writing the JSON-encoded data to the file. It returns an error if any
issues occur during the file creation or encoding process.
***********************************************************************
block.go file :

function HashTXs :
The algorithm is calculating the hash of the transactions of a block 'b'
by concatenating the hashes of its individual transactions and then 
applying the SHA-256 hashing algorithm to the concatenated data.

function NewBlock :
The NewBlock function creates a block using the provided parameters. 
If mine is true, the block is mined using the Mine function with zeroBits 
as a criterion, ensuring that the hash starts with zerosBits. 
The block's nonce is initialized to 0 and incremented until a miner 
constructs the compliant hash, marking the end of the process. 
Otherwise, the hash of the new block is nil, and the nonce is set to 0
until a miner constructs this hash.

function NewGBlock :
The `NewGBlock` function creates the first block (genesis block) of a 
blockchain using a list of Coinbase transactions. It ensures that the 
block's hash starts with a specified number of zeros by calling the 
'Mine' function with the 'zeroBits' parameter. The resulting block is 
then returned.

function IsCorrectlyHashed :
The function checks if a block's hash starts with a specified number of
leading zero bits. It calculates the hash and compares it to a target 
value based on the required zero bits. The function returns 'true' if 
the hash meets the criteria and 'false' otherwise. This is commonly 
used to verify the correctness of mined blocks in a blockchain.

function computeHash :
The 'computeHash' algorithm calculates the hash of a block by 
concatenating and hashing various block details, including the hash 
of the previous block, the hash of transactions, the timestamp, and 
the nonce. The function uses SHA-256 and returns the result as a byte 
array.

function SetHash :

The SetHash function computes the hash of a block using the private 
'computeHash' function and sets the result as the hash of the block.

function Mine :
The 'Mine' function iteratively adjusts the nonce to find a hash for 
the block that meets the specified criteria of starting with a certain
number of leading zeros 'zeroBits'. It computes the hash and updates 
the block's hash until the condition is satisfied.
***********************************************************************
blockchain.go file

function AddBlock :
The idea of ​​the algorithm is to add a new block to the blockchain using
the 'NewBlock' function with the transaction parameters, prevBlock 
which will be an empty byte if it is the 1st block otherwise the hash 
of the previous block, true to mine the block before its addition and 
initialize zeroBits to 16 (choice) to pass it to the Mine function 
called inside the NewBlock function.

function NewTransfertTX :
The NewTransfertTX method creates a transaction to transfer a specified
amount from one address (sender) to another (receiver). It utilizes the
"GetBalance" method to identify the balance of the sender.
If the sender has the required amount, the method constructs a transaction
input using the accumulated value, an output for the receiver with the
specified amount, and an additional output for the remaining balance. 
The transaction is then created and returned.

function NewBlockchain :
Hypothesis :
"CreateBlockchain creates a new blockchain" == The block to create is the genisis block
For each address, a coinbase transaction is created with a reward of 10
(initialized in the 'transaction.go' file). These transactions are 
added to a transaction array. Subsequently, a new block is generated
by passing the transactions and a 'zeroBits' parameter to the 
'NewGBlock' function, as it represents the genesis block of the 
blockchain. This block is added to the blockchain using 'NewBlockchainFromGB',
and the resulting blockchain is returned.

function NewBlockchainFromGB :
The function 'NewBlockchainFromGB' generates a new blockchain using a
valid genesis block. It initializes the blockchain with the genesis 
block's hash and sets the chain to include only the genesis block. 
The resulting blockchain is then returned.

function GetBalance :
The 'GetBalance' method calculates the balance of a given address by 
iterating through all blocks in the blockchain, examining each 
transaction's outputs (TxOuts) and inputs (TxIns). 
It increments the balance for each unspent output with a matching address
and decrements the balance for each spent input with a matching address.
The final balance reflects the total amount available for spending 
associated with the given address.

references :
https://pkg.go.dev/encoding/binary
https://pkg.go.dev/bytes
https://pkg.go.dev/encoding/gob
https://www.youtube.com/watch?v=oCm46sUILcs&list=PL0xRBLFXXsP6-hxQmCDcl_BHJMm0mhxx7
https://blog.logrocket.com/build-blockchain-with-go/
...
While GPT chat is not a formal reference, it's worth mentioning that 
it has played a significant role in helping me articulate the ideas 
behind the algorithms I've developed in a concise manner.
It has also assisted me in commenting the code effectively.














