import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13

import "../."
import "."

FocusScope {
    id: control

    property alias text: textField.text
    property alias placeholderText: textField.placeholderText
    property alias labelText: label.text
    property alias hintText: hint.text
    property alias validator: textField.validator
    property int labelLeftMargin: 5

    implicitWidth: 400
    height: col.height

    onWidthChanged: textField.width = width

    Column {
        id: col
        spacing: 5
        RowLayout {

            width: textField.width
            visible: labelText.length > 0
            
            Label {
                id: label

                leftPadding: labelLeftMargin
            }

            Label {
                id: hint
                Layout.alignment: Qt.AlignRight
                Layout.rightMargin: 5

                color: Qt.darker(Style.textColor3, 1.6)
            }
        }

        TextField {
            id: textField

            color: Style.textColor
            placeholderTextColor: Style.borderColor
            selectByMouse: true
            selectionColor: Style.accentColor2

            background: Rectangle {
                implicitHeight: 40
                implicitWidth: 400
                color: Style.bgInputColor
                border.color: Style.borderColor
            }
        }
    }
}

