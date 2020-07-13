import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Item {
    id: root

    readonly property var wc: mc.workspaceCtrl
    readonly property var ic: wc.inputCtrl

    implicitHeight: 40
    
    Row {
        anchors.left: parent.left
        spacing: 5
        ComboBox {
            id: cbServiceList

            color: Style.primaryColor
            currentIndex: 0

            model: ic.serviceListModel
            onActivated: {
                ic.serviceChanged(displayText)
                cbMethodList.currentIndex = 0
            }
        }
        ComboBox {
            id: cbMethodList

            color: Style.primaryColor
            currentIndex: 0

            model: ic.methodListModel
            onActivated: ic.methodChanged(cbServiceList.displayText, displayText)
        }
    }

    Button {
        id: btnSend
        anchors.right: parent.right
        text: qsTr("Send")
        color: Style.primaryColor
        onClicked: wc.send(cbServiceList.displayText, cbMethodList.displayText)
    } 
}
