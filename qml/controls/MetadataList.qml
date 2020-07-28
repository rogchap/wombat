import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13

import "../."
import "../controls"

Item {
    id: control

    property string titleText
    property string metadata

    implicitHeight: parent.height
    implicitWidth: parent.width

    Column {
        spacing: 5

        Label {
            id: title
            text: titleText
            font.weight: Font.DemiBold
        }
    
        ScrollView {

            height: control.height - title.height - spacing
            width: control.width + 10
            clip: true


            TextEdit {
                id: te

                function getCSS() {
                    return `
                    <style>
                    span {
                        color: ${Style.primaryColor}
                    }
                    </style>
                    `
                }
                text: getCSS() + metadata
                textFormat: TextEdit.RichText
                color: Style.textColor
                selectionColor: Style.accentColor2
                readOnly: true
                selectByMouse: true
                MouseArea {
                    width: parent.width
                    height: parent.height
                    cursorShape: Qt.IBeamCursor
                    enabled: false
                }
            }
        }
    }
}
