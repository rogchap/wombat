import QtQuick 2.13
import QtQuick.Controls 2.13

import "../controls"

ScrollView {

    property var inputData

    padding: 15
    clip: true

    Column {

        spacing: 10

        Label {
            text: inputData.label
        }
    
        MessageFields {
            model: inputData
        }

    }

}
