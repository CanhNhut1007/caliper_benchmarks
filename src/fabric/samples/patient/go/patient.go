package main

import (
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "github.com/hyperledger/fabric-chaincode-go/shim"
        "github.com/hyperledger/fabric/common/util"
        pb "github.com/hyperledger/fabric-protos-go/peer"
)

type ContactInfo struct {
        Relationship    string `json:"relationship"`
        Name            string `json:"name"`
        PhoneNumber     string `json:"phonenumber"`
}

type DateTime struct {
        Date                    uint64 `json:"date"`
        Month                   uint64 `json:"month"`
        Year                    uint64 `json:"year"`
}

type Patient struct {
        ResourceType    string          `json:"resourcetype"`
        Email           string          `json:"email"`
        Name            string          `json:"name"`
        Active          bool            `json:"active"`
        PhoneNumber     string          `json:"phonenumber"`
        Gender          string          `json:"gender"`
        Address         string          `json:"address"`
        DateofBirth     DateTime        `json:"dateofbirth"`
        Photo           string          `json:"photo"`
        Contact         ContactInfo     `json:"contact"`
        HealthRecordID  string          `json:"healthrecordid"`
}

func (t *Patient) Init(stub shim.ChaincodeStubInterface) pb.Response {
        return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. patient
func (t *Patient) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
        // Extract the function and args from the transaction proposal
        function, args := stub.GetFunctionAndParameters()
        fmt.Println("invoke is running " + function)
        var err error
        // Route to the appropriate handler function to interact with the ledger
        if function == "getPatient" {
                return t.getPatient(stub, args)
        } else if function == "createPatient" {
                return t.createPatient(stub, args)
        } else if function == "updatePatient" {
                return t.updatePatient(stub, args)
        } else if function == "queryallPatient" {
                return t.updatePatient(stub, args)
        }

        if err != nil {
                return shim.Error(err.Error())
        }

        // Return the result as success payload
        fmt.Println("invoke did not find func: " + function)
        //error
        return shim.Error("Received unknown function invocation")
}

// Creates patient
func (t *Patient) createPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        // Must have: 0) Email, 1) Name, 2) PhoneNumber, 3) Gender, 4) Address , 5) DayofBirth, 6) Photo, 7)Contact 8) HealthRecordID
    if len(args) != 8 {
                return shim.Error("Incorrect number of arguments. Expecting 9")
        }
        fmt.Println("Arg0:" + args[0]) // Email: nhutori1@gmail.com
        fmt.Println("Arg1:" + args[1]) // Thach Canh Nhut
        fmt.Println("Arg2:" + args[2]) // Phone Number: 0967072612
        fmt.Println("Arg3:" + args[3]) // Fermale
        fmt.Println("Arg4:" + args[4]) // Address
        fmt.Println("Arg5:" + args[5]) // Dayofbirth: 10-07-1998
        fmt.Println("Arg6:" + args[6]) // Photo: Base64(photo)
        fmt.Println("Arg7:" + args[7]) // Contact: (Mother-Tra Thi May-0763915396)

          
        email := args[0]
          // Get the state from the ledger
        existpatient, err := stub.GetState(email)
  
        if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + email + "\"}"
                return shim.Error(jsonResp)
        }
  
        if existpatient != nil {
                jsonResp := "{\"Error\":\"Patient exist\"}"
                return shim.Error(jsonResp)
        }


        dayofbirtharray := strings.Split(args[5], "-")
        date, err  := strconv.ParseUint(dayofbirtharray[0], 10, 64)
        month, err  := strconv.ParseUint(dayofbirtharray[1], 10, 64)
        year, err  := strconv.ParseUint(dayofbirtharray[2], 10, 64)
        fmt.Println(err)
        var dayofbirth = DateTime{ Date: date, Month: month, Year: year}

        contactinfo := strings.Split(args[7], "-")
        var contact = ContactInfo {Relationship : contactinfo[0], Name : contactinfo[1], PhoneNumber : contactinfo[2] }

        chainCodeArgs := util.ToChaincodeArgs("createHealthRecord", args[0])

        response := stub.InvokeChaincode("healthrecord", chainCodeArgs, "hospitalchannel")

        // sDec, _ := b64.StdEncoding.DecodeString(string(response.Payload))
        array := strings.Split(string(response.Payload), ",")
        array = strings.Split(string(array[1]), ":")
        healthrecordid := strings.Trim(array[1], "\"") 

        //jsonresponse, _ := json.Marshal(response)
           
        if response.Status != shim.OK {
                return shim.Error(response.Message)
        }

        // Set args for the created record
        var patient = Patient{ResourceType: "Patient", Email: args[0], Name: args[1], Active: true, PhoneNumber: args[2], Gender: args[3], Address: args[4], DateofBirth : dayofbirth, Photo: args[6],Contact:contact,HealthRecordID:healthrecordid}
   
        patientAsBytes, _ := json.Marshal(patient)

        stub.PutState(args[0], patientAsBytes)

        return shim.Success([]byte(string(patientAsBytes)))
}

func (t *Patient) updatePatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        // Must have: 0) Name, 1) PhoneNumber, 2) Address , 3) DayofBirth, 4) Photo, 5)Contact 
        if len(args) != 6 {
           return shim.Error("Incorrect number of arguments. Expecting 6")
        }
        fmt.Println("Arg0:" + args[0]) // Email: nhutori1@gmail.com
        fmt.Println("Arg1:" + args[1]) // Phone Number: 0967072612
        fmt.Println("Arg2:" + args[2]) // Address
        fmt.Println("Arg3:" + args[3]) // Dateofbirth: 10-07-1998
        fmt.Println("Arg4:" + args[4]) // Photo: Base64(photo)
        fmt.Println("Arg5:" + args[5]) // Contact: (Mother-Tra Thi May-0763915396)
   
        dayofbirtharray := strings.Split(args[3], "-")
           
        date, err  := strconv.ParseUint(dayofbirtharray[0], 10, 64)
        month, err  := strconv.ParseUint(dayofbirtharray[1], 10, 64)
        year, err  := strconv.ParseUint(dayofbirtharray[2], 10, 64)
        fmt.Println(err)
        var dayofbirth = DateTime{ Date: date, Month: month, Year: year}

        contactinfo := strings.Split(args[5], "-")
        var contact = ContactInfo {Relationship : contactinfo[0], Name : contactinfo[1], PhoneNumber : contactinfo[2] }

        email := args[0]
        // get the state information
        bytes, _ := stub.GetState(email)

        if bytes == nil {
                        return shim.Error("Provided email not found!!!")
        }
        // unmarshall the data
        // Read the JSON and Initialize the struct
        var patient  Patient

        _ = json.Unmarshal(bytes, &patient)

        patient.PhoneNumber = string(args[1])
        patient.Address = string(args[2])
        patient.DateofBirth = dayofbirth
        patient.Photo = string(args[4])
        patient.Contact = contact

        jsonpatient, _ := json.Marshal(patient)

        stub.PutState(email, jsonpatient)

        return shim.Success([]byte("Update success!!! "+ string(jsonpatient)))
}

// Retrieve medical record of person with personal number as identifier
func (t *Patient) getPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        var email string // Entity
        var err error

        if len(args) != 1 {
                        return shim.Error("Incorrect number of arguments. Expecting email of the person to query")
        }

        email = args[0]
        // Get the state from the ledger
        patient, err := stub.GetState(email)

        if err != nil {
                        jsonResp := "{\"Error\":\"Failed to get state for " + email + "\"}"
                        return shim.Error(jsonResp)
        }

        if patient == nil {
                        jsonResp := "{\"Error\":\"No medical record with email " + email + "\"}"
                        return shim.Error(jsonResp)
        }

        jsonResp := "{\"Email\":\"" + email + "\",\"Patient\":\"" + string(patient) + "\"}"

        fmt.Printf("Query Response:%s\n", jsonResp)

        return shim.Success(patient)
}

func (t *Patient) queryallPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        if len(args) != 2 {
                return shim.Error("Incorrect number of arguments. ")
        }
	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllPatients:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {
        // Create a new Medical Record
        err := shim.Start(new(Patient))
        if err != nil {
                        fmt.Printf("Error creating new Simple Asset: %s", err)
        }
}