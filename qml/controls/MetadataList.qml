import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13

import "../."
import "../controls"

Item {
    id: control

    property string titleText
    property alias model: lv.model

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

            ListView {
                id: lv

                height: parent.height
                width: parent.width

                delegate: TextEdit {
                    textFormat: TextEdit.RichText
                    leftPadding: 5
                    color: Style.textColor
                    selectionColor: Style.accentColor2
                    text: "<span style=color:" + Style.accentColor + ">" + display + ":</span>&nbsp;<span>" + val + "</span>"
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
}
