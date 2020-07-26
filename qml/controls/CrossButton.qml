import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

AbstractButton {
    id: control

    property color color
    property alias rotation: canvas.rotation

    focusPolicy: Qt.NoFocus

    height: 16
    width: 16

    Canvas {
        id: canvas
        x: 2
        y: 2
        width: 12
        height: 12
        contextType: "2d"

        onPaint: {
            context.reset()
            context.lineWidth = 2
            context.strokeStyle = color
            context.moveTo(0, height / 2);
            context.lineTo(width, height/2);
            context.moveTo(width / 2, 0)
            context.lineTo(width / 2, height)
            context.stroke();
        }
    }

    background: Rectangle {
        color: control.down ? Style.bgColor2 : Style.bgColor
    }
}
