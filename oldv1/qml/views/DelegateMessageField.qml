import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Item {
    id: root

    property alias labelColor: msgLabel.color
    property int labelLeftMargin: 5
    property variant msgOverride

    height: msgPane.height + msgLabel.height + 10
    implicitWidth: msgPane.width

    Label {
        id: msgLabel
        text: label
        anchors.left: parent.left
        anchors.leftMargin: labelLeftMargin
    }

    Label {
        id: msgType
        anchors.left: msgLabel.right
        anchors.leftMargin: 10
        color: Qt.darker(Style.textColor3, 1.6)
        text: message.label
    }

    CheckBox{
        id: enabled
        anchors.left: msgType.right
        anchors.leftMargin: 10
        anchors.verticalCenter: msgType.verticalCenter
        checked: (msgOverride || message).enabled
        onCheckedChanged: (msgOverride || message).enabled = checked 
    }

    Pane {
        id: msgPane

        visible: enabled.checked

        implicitHeight: visible ? msgLoader.height : 0
        anchors.top: msgLabel.bottom

        Loader {
            id: msgLoader

            source: "MessageFields.qml"
            onLoaded: {
                item.model = msgOverride || message
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
