package SAFARIS

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Hash         string
	PreviousHash string
	Ride         Ride
}

type Ride struct {
	DriverID      string
	User          User
	CustomerBid   float64
	DriveBid      float64
	DepartureTime time.Time
	ArrivalTime   time.Time
	TimeTaken     time.Duration
}

type Driver struct {
	Name       string
	ID         string
	VehicleReg string
	// Picture      []byte
	PhoneNumber  string
	TimeStamp    time.Time
	Hash         string
	PreviousHash string
}

type User struct {
	Name        string
	PhoneNumber string
}

type DriverBlock struct {
	Drivers []*Driver
}

type BlockChain struct {
	Blocks []Block
}

// Create a new driver
func (d *DriverBlock) AddDriver(name string, id, regNumber string, num string) {
	previousBlock := d.Drivers[len(d.Drivers)-1]
	previousHash := previousBlock.Hash
	driver := Driver{
		Name:       name,
		ID:         id,
		VehicleReg: regNumber,
		// Picture:      picture,
		PhoneNumber:  num,
		TimeStamp:    time.Now(),
		PreviousHash: previousHash,
	}
	driver.Hash = calculateHash(driver.Name, driver.ID, driver.VehicleReg, driver.PhoneNumber, driver.PreviousHash)
	d.Drivers = append(d.Drivers, &driver)
}

// Create a new ride and add it to a block
func AddRide(driverID string, user User, customerBid, driveBid float64, departureTime, arrivalTime time.Time) Block {
	ride := Ride{
		DriverID:      driverID,
		User:          user,
		CustomerBid:   customerBid,
		DriveBid:      driveBid,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		TimeTaken:     arrivalTime.Sub(departureTime),
	}

	block := Block{
		Ride: ride,
	}
	block.Hash = calculateHash(driverID, user.Name, user.PhoneNumber, fmt.Sprintf("%f", customerBid), fmt.Sprintf("%f", driveBid), departureTime.String(), arrivalTime.String(), fmt.Sprintf("%d", ride.TimeTaken.Milliseconds()))
	return block
}

// Calculate hash based on concatenated parts
func calculateHash(parts ...string) string {
	var buffer bytes.Buffer
	for _, part := range parts {
		buffer.WriteString(part)
	}
	hash := sha256.Sum256(buffer.Bytes())
	return fmt.Sprintf("%x", hash)
}

// Add a block to the blockchain
func (bc *BlockChain) AddBlock(block Block) {
	block.PreviousHash = bc.getLatestBlockHash()
	block.Hash = calculateHash(block.Ride.DriverID, block.Ride.User.Name, block.Ride.User.PhoneNumber, fmt.Sprintf("%f", block.Ride.CustomerBid), fmt.Sprintf("%f", block.Ride.DriveBid), block.Ride.DepartureTime.String(), block.Ride.ArrivalTime.String(), fmt.Sprintf("%d", block.Ride.TimeTaken.Milliseconds()), block.PreviousHash)
	bc.Blocks = append(bc.Blocks, block)
}

// Get the hash of the latest block in the blockchain
func (bc *BlockChain) getLatestBlockHash() string {
	if len(bc.Blocks) == 0 {
		return ""
	}
	return bc.Blocks[len(bc.Blocks)-1].Hash
}
