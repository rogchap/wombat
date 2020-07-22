import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

ScrollView {
    id: root

    padding: 15

    readonly property var oc: mc.workspaceCtrl.outputCtrl

    clip: true
    
    TextArea {
        id: outTxt
        text: oc.output
        color: Style.purpleColor
        readOnly: true
        selectByMouse: true
        selectionColor: Style.accentColor2
        wrapMode: TextEdit.Wrap

        MouseArea {
            width: parent.width
            height: parent.height
            cursorShape: Qt.IBeamCursor
            enabled: false
        }

    }

}
