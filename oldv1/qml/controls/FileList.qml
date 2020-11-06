import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

Item {
    id: control

    property alias labelText: label.text
    property alias actionButtonColor: btnAction.color
    property alias actionButtonText: btnAction.text
    property alias model: listView.model

    signal opened
    signal cleared

    
    implicitWidth: 400
    implicitHeight: 150 + label.height + rowAction.height 

    Label {
        id: label

        anchors {
            left: control.left
            leftMargin: 5
        }
    }

    ScrollView {
        id: scrollView

        height: control.height - label.height - rowAction.height - 10
        width: control.width

        anchors {
            top: label.bottom
            topMargin: 5
        }

        background: Rectangle {
            color: Style.bgInputColor
            border.color: Style.borderColor
        }

        ListView {
            id: listView

            anchors {
                fill: parent
                margins: 5
            }

            clip: true

            delegate: Label {
                width: parent.width

                text: display
                elide: Text.ElideLeft
            }
        }
    }

    Row {
        id: rowAction
        layoutDirection: Qt.RightToLeft
        
        anchors {
            top: scrollView.bottom
            right: scrollView.right
            topMargin: 5
        }

        Button {
            id: btnAction
            hideBorder: false
            text: qsTr("Open")
            onClicked: control.opened()
        }

        Button {
            text: qsTr("Clear")
            onClicked: control.cleared()
        }
    }
}
