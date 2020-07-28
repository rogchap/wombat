import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

ScrollView {
    id: root

    padding: 15

    readonly property var oc: mc.workspaceCtrl.outputCtrl

    clip: true

    function getCSS() {
        return `
        <style>
        .yellow {
            color: ${Style.yellowColor}
        }
        .green {
            color: ${Style.greenColor}
        }
        .red {
            color: ${Style.redColor}
        }
        </style>
        `
    }
    
    TextEdit {
        id: outTxt
        text: getCSS() + oc.stats
        color: Style.textColor3
        readOnly: true
        textFormat: TextEdit.RichText
        selectByMouse: true
        selectionColor: Style.accentColor2

        Component.onCompleted: wrapMode = TextEdit.Wrap

        MouseArea {
            width: parent.width
            height: parent.height
            cursorShape: Qt.IBeamCursor
            enabled: false
        }
    }
}
