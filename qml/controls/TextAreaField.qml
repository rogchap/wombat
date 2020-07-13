import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

Item {
    id: control

    property alias text: textField.text
    property alias placeholderText: textField.placeholderText
    property alias labelText: label.text
    property alias hintText: hint.text

    height: label.height + 150 + 5
    implicitWidth: textField.width

    Label {
        id: label

        anchors {
            left: control.left
            leftMargin: 5
        }
    }

    Label {
        id: hint
        anchors {
            right: control.right
            rightMargin: 5
        }

        color: Qt.darker(Style.textColor3, 1.6)
    }

    ScrollView {
        implicitHeight: 150
        implicitWidth: 400

        anchors {
            top: label.bottom
            topMargin: 5
        }
            background: Rectangle {
                color: Style.bgInputColor
                border.color: Style.borderColor
            }

        TextArea {
            id: textField
            color: Style.textColor
            placeholderTextColor: Style.borderColor
            selectByMouse: true
            selectionColor: Style.accentColor2
        }
    }
}


