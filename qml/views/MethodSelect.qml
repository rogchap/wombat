import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {
    id: root
    padding: 0
    implicitHeight: 40
    
    Row {
        anchors.left: parent.left
        spacing: 5
        ComboBox {
            color: Style.primaryColor
            model: ["s12.Authorization.Access", "Second", "Third"]
        }
        ComboBox {
            color: Style.primaryColor
            model: ["GetAccessToken", "Second", "Third"]
        }
    }

    Button {
        id: btnSend
        anchors.right: parent.right
        text: qsTr("Send")
        color: Style.primaryColor
    }
}
