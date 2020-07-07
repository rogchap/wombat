import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

Item {
    id: control

    property alias text: textField.text
    property alias placeholderText: textField.placeholderText
    property alias labelText: label.text

    height: label.height + textField.height + 5
    implicitWidth: 200

    Label {
        id: label

        anchors {
            left: control.left
            leftMargin: 5
        }
    }

    TextField {
        id: textField

        anchors {
            top: label.bottom
            topMargin: 5
        }

        color: Style.textColor
        placeholderTextColor: Style.borderColor

        background: Rectangle {
            implicitHeight: 40
            implicitWidth: 400
            color: Style.bgInputColor
            border.color: Style.borderColor
        }
    }
}

