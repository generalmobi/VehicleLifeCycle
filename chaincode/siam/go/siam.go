/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
         "bytes"
         "encoding/json"
         "fmt"
         "strconv"
         "time"
         "github.com/hyperledger/fabric/core/chaincode/shim"
         sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }

 type Registration struct {
        RegistrationDate   string `json:registrationDate` 
        Validity   string `json:validty`
        ChasisNumber string  `json:chasis`
        RegistrationNumber string `json:registrationNumber`
        Owner string `json:owner`
}


 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Car struct {
         Make   string `json:"make"`
         Model  string `json:"model"`
         Color  string `json:"color"`
         Cc     string `json:cc`
         ChasisNumber string  `json:chasisNumber`   
 }
 
 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
         return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
         // Retrieve the requested Smart Contract function and arguments
         function, args := APIstub.GetFunctionAndParameters()
         // Route to the appropriate handler function to interact with the ledger appropriately
         if function == "registerCar" {
                 return s.registerCar(APIstub, args)
         } else if function == "createCar" {
                 return s.createCar(APIstub, args)
         } else if function == "getCarHistory" {
                 return s.getCarHistory(APIstub,args)
         } else if function == "changeCarOwner" {
                 return s.changeCarOwner(APIstub, args)
         } else if function == "getCar" {
                return s.getCar(APIstub, args)
         } else if function == "scrapCar" {
                return s.scrapCar(APIstub, args)
         }
         
         return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 1 {
                 return shim.Error("Incorrect number of arguments. Expecting 1")
         }
         var chasisNumber= args[0]
         carAsBytes, _ := APIstub.GetState(chasisNumber)
         return shim.Success(carAsBytes)
 }
  
 func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 5 {
                 return shim.Error("Incorrect number of arguments. Expecting 5")
         }
 
         fmt.Printf("- start createCar for : %s\n", args[0])

         var car = Car{Make: args[1], Model: args[2], Color: args[3], Cc: args[4]}
 
         carAsBytes, _ := json.Marshal(car)
         APIstub.PutState(args[0], carAsBytes)
 
         return shim.Success(nil)
 }
 

 func (s *SmartContract) registerCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
        if len(args) != 4 {
                return shim.Error("Incorrect number of arguments. Expecting 5")
        }

        var registration = Registration{RegistrationDate: args[1], ChasisNumber: args[2], Validity: args[3]}

        registrationAsBytes, _ := json.Marshal(registration)
        APIstub.PutState(args[0], registrationAsBytes)

        return shim.Success(nil)
}

 

func (s *SmartContract) getCarHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) < 1 {
            return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        registrtionNumber := args[0]
    
        fmt.Printf("- start getCarHistory for : %s\n", registrtionNumber)
    
        resultsIterator, err := APIstub.GetHistoryForKey(registrtionNumber)
        if err != nil {
            return shim.Error(err.Error())
        }
        defer resultsIterator.Close()
    
        var buffer bytes.Buffer
        buffer.WriteString("[")
    
        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {
            response, err := resultsIterator.Next()
            if err != nil {
                return shim.Error(err.Error())
            }
            if bArrayMemberAlreadyWritten == true {
                buffer.WriteString(",")
            }
            buffer.WriteString("{\"TxId\":")
            buffer.WriteString("\"")
            buffer.WriteString(response.TxId)
            buffer.WriteString("\"")
    
            buffer.WriteString(", \"Value\":")
            if response.IsDelete {
                buffer.WriteString("null")
            } else {
                buffer.WriteString(string(response.Value))
            }
    
            buffer.WriteString(", \"Timestamp\":")
            buffer.WriteString("\"")
            buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
            buffer.WriteString("\"")
    
            buffer.WriteString(", \"IsDelete\":")
            buffer.WriteString("\"")
            buffer.WriteString(strconv.FormatBool(response.IsDelete))
            buffer.WriteString("\"")
            buffer.WriteString("}")
            bArrayMemberAlreadyWritten = true
        }
        buffer.WriteString("]")
    
        fmt.Printf("- getCarHistory returning:\n%s\n", buffer.String())
    
        return shim.Success(buffer.Bytes())
    
    }




    func (s *SmartContract) getCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 2")
        }

        carAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(carAsBytes)
}

func (s *SmartContract) scrapCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        var jsonResp string
        var registrationJson Registration
        if len(args) != 1 {
            return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        registrationNumber := args[0]
        // to maintain the color~name index, we need to read the marble first and get its color
        valAsbytes, err := APIstub.GetState(registrationNumber) //get the marble from chaincode state
        if err != nil {
            jsonResp = "{\"Error\":\"Failed to get state for " + registrationNumber + "\"}"
            return shim.Error(jsonResp)
        } else if valAsbytes == nil {
            jsonResp = "{\"Error\":\"registration does not exist: " + registrationNumber + "\"}"
            return shim.Error(jsonResp)
        }
        err = json.Unmarshal([]byte(valAsbytes), &registrationJson)
        if err != nil {
            jsonResp = "{\"Error\":\"Failed to decode JSON of: " + registrationNumber + "\"}"
            return shim.Error(jsonResp)
        }
    
        err = APIstub.DelState(registrationNumber) //remove the marble from chaincode state
        if registrationJson.ChasisNumber != ""  {
                err = APIstub.DelState(registrationJson.ChasisNumber) //remove the marble from chaincode state
        }

        if err != nil {
            return shim.Error("Failed to delete state:" + err.Error())
        }
        return shim.Success(nil)
    }




 func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
         if len(args) != 2 {
                 return shim.Error("Incorrect number of arguments. Expecting 2")
         }
 
         registrationAsBytes, _ := APIstub.GetState(args[0])
         registration := Registration{}
 
         json.Unmarshal(registrationAsBytes, &registration)
         registration.Owner = args[1]
 
         registrationAsBytes, _ = json.Marshal(registration)
         APIstub.PutState(args[0], registrationAsBytes)
 
         return shim.Success(nil)
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
         // Create a new Smart Contract
         err := shim.Start(new(SmartContract))
         if err != nil {
                 fmt.Printf("Error creating new Smart Contract: %s", err)
         }
 }