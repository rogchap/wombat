import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

ComboBox {
    id: control

    property alias color: content.color

    textRole: "display"

    delegate: ItemDelegate {
        width: control.width
        contentItem: Text {
            text: control.textRole ? (Array.isArray(control.model) ? modelData[control.textRole] : model[control.textRole]) : modelData
            color: Style.textDarkColor
            font.weight: control.currentIndex === index ? Font.DemiBold : Font.Normal
            elide: Text.ElideLeft
            verticalAlignment: Text.AlignVCenter
        }
        highlighted: control.highlightedIndex === index
    } 

    indicator: Canvas {
        id: canvas
        // x: control.width - width 
        x: control.width - width - control.rightPadding
        y: control.topPadding + (control.availableHeight - height) / 2
        width: 12
        height: 8
        contextType: "2d"

        onPaint: {
            context.reset();
            context.moveTo(0, 0);
            context.lineTo(width, 0);
            context.lineTo(width / 2, height);
            context.closePath();
            context.fillStyle = Style.bgColor3
            context.fill();
        }
    }

    contentItem: Label {
        id: content
        leftPadding: 12
        rightPadding: control.indicator.width + 12

        text: control.displayText
        font: control.font
        color: Style.textColor
        verticalAlignment: Text.AlignVCenter
        elide: Text.ElideLeft

        background: Rectangle {
            color: "transparent"
        }
    }

    background: Rectangle {
        implicitWidth: 120
        implicitHeight: 40
        color: control.down? Style.bgColor2 : Style.bgColor
    }

}
