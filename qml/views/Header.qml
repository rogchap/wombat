import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13

import "../."
import "../controls"

Pane {
    id: root 
    background: Rectangle {
        implicitHeight: 40 + topPadding + bottomPadding
        color: Style.bgColor

        Rectangle {
            color: Style.borderColor
            width: parent.width
            height: 1
            anchors.top: parent.bottom
        }
    }

    Button {
        text: qsTr("Workspace")
        bgColor: Style.accentColor3

        onClicked: wkspOptions.open()

        WorkspaceOptions {
            id: wkspOptions
        }
    }

    Label {
        id: lblAddr
        anchors.centerIn: parent
        text: mc.workspaceCtrl.addr
        font.pointSize: 14
        font.weight: Font.DemiBold
        color: Style.primaryColor
    }

    Label {
        function getColor() {
            switch(text) {
                case "ready":
                    return Style.greenColor
                case "connecting":
                case "idle":
                    return Style.yellowColor
                default:
                    return Style.redColor
            }
        }
        anchors.top: lblAddr.bottom
        anchors.horizontalCenter: lblAddr.horizontalCenter
        text: mc.workspaceCtrl.connState.toLowerCase().replace("_", " ")
        color: getColor()
    }


}

