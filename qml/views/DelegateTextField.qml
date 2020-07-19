import QtQuick 2.13
import QtQuick.Controls 2.13

import "../controls"

TextField {
    id: root

    labelText: label
    hintText: type
    text: val
    
    Component.onCompleted: {
        switch(type) {
            case "double":
            case "float":
            validator = vDouble
            break
            case "int64":
            case "fixed64":
            case "sfixed64":
            case "sint64":
            validator = vInt64
            break
            case "uint64":
            case "uint32":
            validator = vUInt64
            break
            case "int32":
            case "fixed32":
            case "sfixed32":
            case "sint32":
            validator = vInt32
            break
        }
    }

    DoubleValidator { id: vDouble }
    RegExpValidator { id: vInt64; regExp: /^-?\d+/ }
    RegExpValidator { id: vUInt64; regExp: /\d+/ }
    IntValidator { id: vInt32; bottom: -2147483648; top: 2147483647 }

}

