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
            text: qsTr("New Workspace")
            color: Style.accentColor3
        }

        Label {
            anchors.right: parent.right
            height: parent.height
            verticalAlignment: Text.AlignVCenter
            text: qsTr("connected")
            color: "green"
        }

    
}

