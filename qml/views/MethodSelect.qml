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
            id: cbServiceList

            color: Style.primaryColor
            currentIndex: 0

            model: mc.serviceList
            onActivated: {
                mc.serviceChanged(displayText)
                cbMethodList.currentIndex = 0
            }
        }
        ComboBox {
            id: cbMethodList

            color: Style.primaryColor
            currentIndex: 0

            model: mc.methodList
            onActivated: mc.methodChanged(cbServiceList.displayText, displayText)
        }
    }

    Button {
        id: btnSend
        anchors.right: parent.right
        text: qsTr("Send")
        color: Style.primaryColor
    }
}
