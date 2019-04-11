package llrp

// C1G2PC generates C1G2PC parameter from hexpc string.
func C1G2PC(pc uint16) []byte {
	var data = []interface{}{
		uint8(140), // 1+uint7(Type=12)
		pc,         // PC bits
	}
	return Pack(data)
}

// C1G2ReadOpSpecResult generates C1G2ReadOpSpecResult parameter from readData.
func C1G2ReadOpSpecResult(readData []byte) []byte {
	var data = []interface{}{
		uint16(349), // Rsvd+Type=
		uint16(11),  // Length
		uint8(0),    // Result
		uint16(9),   // OpSpecID
		uint16(1),   // ReadDataWordCount
		readData,    // ReadData
	}
	return Pack(data)
}

// ConnectionAttemptEvent generates ConnectionAttemptEvent parameter.
func ConnectionAttemptEvent() []byte {
	var data = []interface{}{
		uint16(256), // Rsvd+Type=256
		uint16(6),   // Length
		uint16(0),   // Status(Success=0)
	}
	return Pack(data)
}

// EPCData generates EPCData parameter from its length and epcLength, and epc.
func EPCData(length uint16, epcLengthBits uint16, epc []byte) []byte {
	var data []interface{}
	if epcLengthBits == 96 {
		data = []interface{}{
			uint8(141), // 1+uint7(Type=13)
			epc,        // 96-bit EPCData string
		}
	} else {
		data = []interface{}{
			uint16(241),           // uint8(0)+uint8(Type=241)
			uint16(length),        // Length
			uint16(epcLengthBits), // EPCLengthBits
			epc, // EPCData string
		}
	}
	return Pack(data)
}

// KeepaliveSpec generates KeepaliveSpec parameter.
func KeepaliveSpec() []byte {
	var data = []interface{}{
		uint16(220),   // Rsvd+Type=220
		uint16(9),     // Length
		uint8(1),      // KeepaliveTriggerType=Periodic(1)
		uint32(10000), // TimeInterval=10000
	}
	return Pack(data)
}

// Status generates LLRPStatus parameter.
func Status() []byte {
	var data = []interface{}{
		uint16(287), // Rsvd+Type=287
		uint16(8),   // Length
		uint16(0),   // StatusCode=M_Success(0)
		uint16(0),   // ErrorDescriptionByteCount=0
	}
	return Pack(data)
}

// PeakRSSI generates PeakRSSI parameter.
func PeakRSSI() []byte {
	var data = []interface{}{
		uint8(134), // 1+uint7(Type=6)
		uint8(203), // PeakRSSI
	}
	return Pack(data)
}

// ReaderEventNotificationData generates ReaderEventNotification parameter.
func ReaderEventNotificationData(currentTime uint64) []byte {
	utcTimeStamp := UTCTimeStamp(currentTime)
	connectionAttemptEvent := ConnectionAttemptEvent()
	readerEventNotificationDataLength := len(utcTimeStamp) +
		len(connectionAttemptEvent) + 4 // Rsvd+Type+length=32bits=4bytes
	var data = []interface{}{
		uint16(246),                               // Rsvd+Type=246 (ReaderEventNotificationData parameter)
		uint16(readerEventNotificationDataLength), // Length
		utcTimeStamp,
		connectionAttemptEvent,
	}
	return Pack(data)
}

/*
// TagReportData generates TagReportData parameter from epcData, peakRSSI, airProtocolTagData, opSpecResult.
func TagReportData(epcData []byte, airProtocolTagData []byte) []byte {
	tagReportDataLength := len(epcData) + len(airProtocolTagData) +
		4 // Rsvd+Type+length->32bits=4bytes
	var data = []interface{}{
		uint16(240),                 // Rsvd+Type=240 (TagReportData parameter)
		uint16(tagReportDataLength), // Length
		epcData,
		airProtocolTagData,
	}
	return Pack(data)
}
*/

// UTCTimeStamp generates UTCTimeStamp parameter at the current time.
func UTCTimeStamp(currentTime uint64) []byte {
	var data = []interface{}{
		uint16(128), // Rsvd+Type=128
		uint16(12),  // Length
		currentTime, // Microseconds
	}
	return Pack(data)
}

//GeneralDeviceCapabilities : Generates General Device Capabilities
func GeneralDeviceCapabilities() []byte {
	var data = []interface{}{
		uint16(137),                        //Type 137
		uint16(28),                         //Length
		uint16(52),                         //Max Antenna
		uint16(16384),                      //UTC clock support
		uint32(25882),                      //Manufacturer
		uint32(2001007),                    //Model
		[]byte("000a352e31342e302e323430"), //firmware version 5.14.0.240
	}
	return Pack(data)
}

//LlrpCapabilities : generates LLRP_CAPABILITIES
func LlrpCapabilities() []byte {
	var data = []interface{}{
		uint16(142),  //type 142
		uint16(28),   //length
		uint8(72),    //rf survery = no, buffer fille warning = yes, client request opspec = no, tag inventory = no, supoprt event = yes
		uint8(1),     // max priotity level supported
		uint16(0),    //client request opsec timeout
		uint32(1),    // max num of rospec
		uint32(32),   //max num of spec per rospec
		uint32(1),    //max num of inventory spec per AIspec
		uint32(1508), //max num of accessSpec
		uint32(8),    //max num of opspec per AccessSpec
	}
	return Pack(data)
}

//ReguCapabilities : generates Regulatory Capabilities
func ReguCapabilities() []byte {
	var data = []interface{}{
		uint16(143), //type 143
		uint16(8),   //length
		uint16(840), // country code
		uint16(1),   //comm standards, fcc part 15
	}
	return Pack(data)
}

//C1G2llrpCapabilities : Generates C1G2llrpCapabilities
func C1G2llrpCapabilities() []byte {
	var data = []interface{}{
		uint16(327), //type 327
		uint16(7),   //length
		uint8(64),   //some params
		uint16(2),   //max num of selectec filter per query
	}
	return Pack(data)
}

//GetReaderConfigResponseIdentification : Generate Identification
func GetReaderConfigResponseIdentification() []byte {
	var data = []interface{}{
		uint16(218),                  //type
		uint16(15),                   //length
		uint8(0),                     //id type
		[]byte("0008001625ffff11c1"), //Reader ID
	}
	return Pack(data)
}

func AntennaProperties(id uint16) []byte {
	var data = []interface{}{
		uint16(221), //type
		uint16(9),   //length
		uint16(128), //antenna connected
		id,
		uint16(0), //gain
	}
	return Pack(data)
}
