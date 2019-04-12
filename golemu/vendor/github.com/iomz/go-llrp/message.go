package llrp

// Keepalive generates Keepalive message.
func Keepalive(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),              // Length
		messageID,               // ID
	}
	return Pack(data)
}

// KeepaliveAck generates KeepaliveAck message.
func KeepaliveAck(messageID uint32) []byte {
	var data = []interface{}{
		uint16(KeepaliveAckHeader), // Rsvd+Ver+Type=62 (KEEPALIVE)
		uint32(10),                 // Length
		messageID,                  // ID
	}
	return Pack(data)
}

// ReaderEventNotification generates ReaderEventNotification message.
func ReaderEventNotification(messageID uint32, currentTime uint64) []byte {
	readerEventNotificationData := ReaderEventNotificationData(currentTime)
	readerEventNotificationLength :=
		len(readerEventNotificationData) + 10 // Rsvd+Ver+Type+Length+ID->80bits=10bytes
	var data = []interface{}{
		uint16(ReaderEventNotificationHeader), // Rsvd+Ver+Type=63 (READER_EVENT_NOTIFICATION)
		uint32(readerEventNotificationLength), // Length
		messageID, // ID
		readerEventNotificationData,
	}
	return Pack(data)
}

// SetReaderConfig generates SetReaderConfig message.
func SetReaderConfig(messageID uint32) []byte {
	keepaliveSpec := KeepaliveSpec()
	setReaderConfigLength :=
		len(keepaliveSpec) + 11 // Rsvd+Ver+Type+Length+ID+R+Rsvd->88bits=11bytes
	var data = []interface{}{
		uint16(SetReaderConfigHeader), // Rsvd+Ver+Type=3 (SET_READER_CONFIG)
		uint32(setReaderConfigLength), // Length
		messageID,                     // ID
		uint8(0),                      // RestoreFactorySetting(no=0)+Rsvd
		keepaliveSpec,
	}
	return Pack(data)
}

// SetReaderConfigResponse generates SetReaderConfigResponse message.
func SetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	setReaderConfigResponseLength :=
		len(llrpStatus) + 10 // Rsvd+Ver+Type+Length+ID+R+Rsvd->80bits=10bytes
	var data = []interface{}{
		uint16(SetReaderConfigResponseHeader), // Rsvd+Ver+Type=13 (SET_READER_CONFIG_RESPONSE)
		uint32(setReaderConfigResponseLength), // Length
		messageID, // ID
		llrpStatus,
	}
	return Pack(data)
}

//GetReaderCapability :
func GetReaderCapability(messageID uint32) []byte {
	getReaderCapabilityLength := 1 + 10
	var data = []interface{}{
		uint16(GetReaderCapabilityHeader),
		uint32(getReaderCapabilityLength),
		messageID,
		uint8(0), //all capabilities
	}
	return Pack(data)
}

//GetReaderCapabilityResponse :
func GetReaderCapabilityResponse(messageID uint32) []byte {
	getReaderCapabilityResponseLength := 89 //2+4+4+8(llrpstatus)+28+28+8+7
	llrpStatus := Status()
	generalCapabilites := GeneralDeviceCapabilities()
	llrpCapabilities := LlrpCapabilities()
	//c1g2llrpCapabilities := C1G2llrpCapabilities()
	reguCapabilitles := ReguCapabilities()
	var data = []interface{}{
		uint8(4),  //version
		uint8(11), //type
		uint32(getReaderCapabilityResponseLength),
		messageID,
		llrpStatus,
		generalCapabilites,
		llrpCapabilities,
		reguCapabilitles,
		uint8(0),
		uint8(0),
		uint8(0),
		//c1g2llrpCapabilities,
	}
	return Pack(data)
}

//GetReaderConfigResponse :
func GetReaderConfigResponse(messageID uint32) []byte {
	llrpStatus := Status()
	identification := GetReaderConfigResponseIdentification()
	length := 1 + 1 + 2 + 8 + 15 + 9
	var data = []interface{}{
		uint8(4),  //version 1.0.1
		uint8(12), //type
		uint32(length),
		messageID,
		llrpStatus,
		identification,
		AntennaProperties(1),
	}

	return Pack(data)
}

//DeleteAccessSpecResponse : Delete Access Spec Response
func DeleteAccessSpecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint8(4),   //version 1.0.1
		uint8(51),  //type
		uint16(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//DeleteRospecResponse : Delete RoSpec Response
func DeleteRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint8(4),   //version 1.0.1
		uint8(31),  //type
		uint16(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//AddRospecResponse : Add ROSpec Response
func AddRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint8(4),   //version 1.0.1
		uint8(30),  //type
		uint16(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//EnableRospecResponse : Enabled Rospec Response
func EnableRospecResponse(messageID uint32) []byte {
	llrpStatus := Status()
	var data = []interface{}{
		uint8(4),   //version 1.0.1
		uint8(34),  //type
		uint16(18), //length
		messageID,
		llrpStatus,
	}
	return Pack(data)
}

//ReceiveSensitivityEntries : Generates ReceiveSensitivityEntries used in General capabilities
func ReceiveSensitivityEntries(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := ReceiveSensitivityEntry(uint16(i))
		data = append(data, x)
	}
	return data
}

//ReceiveSensitivityEntry :
func ReceiveSensitivityEntry(id uint16) []interface{} {
	var data = []interface{}{
		uint16(139), //type
		uint16(8),   //length
		uint16(id),  //length
		uint16(11),  //receive sentitvitiy value
	}
	return data
}

//GPIOCapabilities : Generates GPIO capabilities proeprty
func GPIOCapabilities() []byte {
	var data = []interface{}{
		uint16(141), //type
		uint16(8),   //length
		uint16(0),   //num of GPI port
		uint16(0),   //num of GPO port
	}
	return Pack(data)
}

//AntennaAirPortList :
func AntennaAirPortList(numOfAntennas int) []interface{} {
	var data = []interface{}{}
	for i := 1; i <= numOfAntennas; i++ {
		x := AntennaAirPort(uint16(i))
		data = append(data, x)
	}
	return data
}

//AntennaAirPort :
func AntennaAirPort(id uint16) []interface{} {
	var data = []interface{}{
		uint16(140), //type
		uint16(9),   //length
		id,
		uint16(1), //num of protocols
		uint8(1),  //protocol id : EPCGlobal Class 1 Gen 2
	}
	return data
}
