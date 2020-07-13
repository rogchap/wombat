import QtQuick 2.13
import QtQuick.Controls 2.13

import "../controls"

TextField {
    labelText: label
    hintText: type.substring(5, type.length).toLowerCase()
    text: val
    onTextChanged: root.model.updateFieldValue(index, text) 
}

