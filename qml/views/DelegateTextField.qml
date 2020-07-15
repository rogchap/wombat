import QtQuick 2.13
import QtQuick.Controls 2.13

import "../controls"

TextField {
    labelText: label
    hintText: type
    text: val
    onTextChanged: root.model.updateFieldValue(index, text) 
    
    Component.onCompleted: {
        switch(type) {
            case "TYPE_DOUBLE":
            case "TYPE_FLOAT":
            validator = vDouble
            break
            case "TYPE_INT64":
            case "TYPE_FIXED64":
            case "TYPE_SFIXED64":
            case "TYPE_SINT64":
            validator = vInt64
            break
            case "TYPE_UINT64":
            case "TYPE_UINT32":
            validator = vUInt64
            break
            case "TYPE_INT32":
            case "TYPE_FIXED32":
            case "TYPE_SFIXED32":
            case "TYPE_SINT32":
            validator = vInt32
            break
        }
    }

    DoubleValidator { id: vDouble }
    RegExpValidator { id: vInt64; regExp: /^-?\d+/ }
    RegExpValidator { id: vUInt64; regExp: /\d+/ }
    IntValidator { id: vInt32; bottom: -2147483648; top: 2147483647 }

}

