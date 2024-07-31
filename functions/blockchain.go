package SAFARIS

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Ride struct {
	// DriverID      string
	// User          User
	Hash          string
	PreviousHash  string
	CustomerBid   float64
	DriveBid      float64
	DepartureTime time.Time
	ArrivalTime   time.Time
	// TimeTaken     time.Duration
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
	Blocks []Ride
}

// Create a new driver
func (d *DriverBlock) AddDriver(name, id, regNumber, num string) {
	fmt.Println(name)
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
	driver.Hash = CalculateHash(driver.Name, driver.ID, driver.VehicleReg, driver.PhoneNumber, driver.PreviousHash)
	d.Drivers = append(d.Drivers, &driver)
}

// Create a new ride and add it to a block
func AddRide(customerBid, driveBid float64, departureTime, arrivalTime time.Time) Ride {
	ride := Ride{
		CustomerBid:   customerBid,
		DriveBid:      driveBid,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
	}

	ride.Hash = CalculateHash(fmt.Sprintf("%f", customerBid), fmt.Sprintf("%f", driveBid), departureTime.String(), arrivalTime.String())
	return ride
}

// Calculate hash based on concatenated parts
func CalculateHash(parts ...string) string {
	var buffer bytes.Buffer
	for _, part := range parts {
		buffer.WriteString(part)
	}
	hash := sha256.Sum256(buffer.Bytes())
	return fmt.Sprintf("%x", hash)
}

// Add a block to the blockchain
func (bc *BlockChain) AddRide(block Ride) {
	block.PreviousHash = bc.getLatestBlockHash()
	block.Hash = CalculateHash(fmt.Sprintf("%f", block.CustomerBid), fmt.Sprintf("%f", block.DriveBid), block.DepartureTime.String(), block.ArrivalTime.String(), block.PreviousHash)
	bc.Blocks = append(bc.Blocks, block)
}

// Get the hash of the latest block in the blockchain
func (bc *BlockChain) getLatestBlockHash() string {
	if len(bc.Blocks) == 0 {
		return ""
	}
	return bc.Blocks[len(bc.Blocks)-1].Hash
}
