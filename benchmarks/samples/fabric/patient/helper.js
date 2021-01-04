/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';

let numbers = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
let genders = ['Fermale', 'Male'];
let names = ['Tomoko', 'Brad', 'Jin Soo', 'Max', 'Adrianna', 'Michel', 'Aarav', 'Pari', 'Valeria', 'Shotaro'];
let addresses = ['120 Street', 'Tran Hung Dao Street', 'Quang Trung Street', 'Ho Chi Minh city', 'Ha Noi', 'Hue', 'Quang Ngai', 'Quang Nam', 'Binh Dinh', 'Ha Tinh'];
let dates = ['1', '2', '3', '4', '5', '6', '7', '8', '9','10','11', '12', '13', '14', '15', '16', '17', '18', '19', '20','21','22','23','24','25','26','27','28'];
let months = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
let years = ['1990', '1991', '1992', '1993', '1994', '1995', '1996', '1997', '1998', '1999'];
let contactnames = ['Tra Thi May','Nguyen Thi Anh','Nguyen Thi Thanh','Nguyen Thi Thuong','Thach Canh Nhut','Thach Canh Thanh','Nguyen Van A','Nguyen Van B','Nguyen Thi C'];
let relationships = ['Friend','Mother','Sister','Brother','Father','GrandMother','GrandFather'];

let email;
let txIndex = 0;

module.exports.createPatient = async function (bc, workerIndex, args) {

    while (txIndex < args.assets) {
        txIndex++;
        let email = 'canhnhutori' + txIndex.toString();
        let name = names[Math.floor(Math.random() * names.length)];
        let phonenumber = '0967' +  numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] +  numbers[Math.floor(Math.random() * numbers.length)] ;
        let gender = genders[Math.floor(Math.random() * genders.length)];
        let address = addresses[Math.floor(Math.random() * addresses.length)];
        let dateofbirth = dates[Math.floor(Math.random() * dates.length)] +'-'+ months[Math.floor(Math.random() * months.length)] +'-'+ years[Math.floor(Math.random() * years.length)];
        let photo ='photo';
        let contact= relationships[Math.floor(Math.random() * relationships.length)] + '-' +contactnames[Math.floor(Math.random() * contactnames.length)] + '-' + '0987654321';

        let myArgs = {
            contractId: 'patient',
            contractVersion: 'v1',
            contractFunction: 'createPatient',
            contractArguments: [email,name,phonenumber,gender,address,dateofbirth,photo,contact],
            timeout: 30
        };

        await bc.sendRequests(myArgs);
    }

};
