
'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');


let numbers = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
let genders = ['Fermale', 'Male'];
let names = ['Tomoko', 'Brad', 'Jin Soo', 'Max', 'Adrianna', 'Michel', 'Aarav', 'Pari', 'Valeria', 'Shotaro'];
let addresses = ['120 Street', 'Tran Hung Dao Street', 'Quang Trung Street', 'Ho Chi Minh city', 'Ha Noi', 'Hue', 'Quang Ngai', 'Quang Nam', 'Binh Dinh', 'Ha Tinh'];
let dates = ['1', '2', '3', '4', '5', '6', '7', '8', '9','10','11', '12', '13', '14', '15', '16', '17', '18', '19', '20','21','22','23','24','25','26','27','28'];
let months = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
let years = ['1990', '1991', '1992', '1993', '1994', '1995', '1996', '1997', '1998', '1999'];
let contactnames = ['Tra Thi May','Nguyen Thi Anh','Nguyen Thi Thanh','Nguyen Thi Thuong','Thach Canh Nhut','Thach Canh Thanh','Nguyen Van A','Nguyen Van B','Nguyen Thi C'];
let relationships = ['Friend','Mother','Sister','Brother','Father','GrandMother','GrandFather'];


/**
 * Workload module for the benchmark round.
 */
class CreatePatientWorkload extends WorkloadModuleBase {
    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.txIndex = 0;
    }

    /**
     * Assemble TXs for the round.
     * @return {Promise<TxStatus[]>}
     */
    async submitTransaction() {
        this.txIndex++;
        let Email = 'anhori' + this.txIndex.toString() + '@gmail.com';
        let Name = names[Math.floor(Math.random() * names.length)];
        let PhoneNumber = '0967' +  numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] + numbers[Math.floor(Math.random() * numbers.length)] +  numbers[Math.floor(Math.random() * numbers.length)] ;
        let Gender = genders[Math.floor(Math.random() * genders.length)];
        let Address = addresses[Math.floor(Math.random() * addresses.length)];
        let Dateofbirth = dates[Math.floor(Math.random() * dates.length)] +'-'+ months[Math.floor(Math.random() * months.length)] +'-'+ years[Math.floor(Math.random() * years.length)];
        let Photo ='';
        let Contact= relationships[Math.floor(Math.random() * relationships.length)] + '-' +contactnames[Math.floor(Math.random() * contactnames.length)] + '-' + phonenumber;

        let args = {
            contractId: 'patient',
            contractVersion: 'v1',
            contractFunction: 'createPatient',
            contractArguments: [Email,Name,PhoneNumber,Gender,Address,Dateofbirth,Photo,Contact],
            timeout: 30
        };

        await this.sutAdapter.sendRequests(args);
    }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
    return new CreatePatientWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
