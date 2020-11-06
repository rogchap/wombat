import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

TabButton {
    id: control

    contentItem: Label {
        padding: 7
        anchors.centerIn: control
        text: control.text
        font.weight: checked ? Font.DemiBold : Font.Normal
        color: checked ? Style.textColor : Style.textColor2
    }

    background: Rectangle {
        width: parent.width
        implicitHeight: 44
        color: Style.bgColor

        Rectangle {
            width: parent.width
            height: 1
            color: Style.borderColor
            anchors.bottom: parent.top
            visible: control.checked ? true : false
        }

        Rectangle {
            width: 1
            height: parent.height
            color: Style.borderColor
            anchors.right: parent.left
            visible: control.checked ? true : false
        }
        Rectangle {
            width: 1
            height: parent.height
            color: Style.borderColor
            anchors.left: parent.right
            visible: control.checked ? true : false
        }
        Rectangle {
            width: parent.width
            height: 1
            color: Style.borderColor
            anchors.bottom: parent.bottom
            visible: control.checked ? false : true
        }
    }
    leftPadding: 15
    rightPadding: 15

    width: implicitWidth
}

