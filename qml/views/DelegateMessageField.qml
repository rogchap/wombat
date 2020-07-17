import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Item {
    height: msgPane.height + msgLabel.height + 10

    Label {
        id: msgLabel
        text: label
        anchors.left: parent.left
        anchors.leftMargin: 5
    }

    Label { 
        anchors.left: msgLabel.right
        anchors.leftMargin: 10
        color: Qt.darker(Style.textColor3, 1.6)
        text: message.label
    }

    Pane {
        id: msgPane

        width: msgLoader.width
        height: msgLoader.height
        anchors.top: msgLabel.bottom

        Loader {
            id: msgLoader

            source: "MessageFields.qml"
            onLoaded: {
                item.model = message
            }
        }

        Rectangle {
            width: 1
            height: msgPane.height + 5
            color: Style.accentColor
            anchors.left: parent.left
            anchors.top: parent.top
            anchors.leftMargin: -7
            anchors.topMargin: -5
        }
    }
}
