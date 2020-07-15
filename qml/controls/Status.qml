import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

Label {

    property int code: 99

    topPadding: 5
    bottomPadding: 5
    leftPadding: 8
    rightPadding: 8

    background: Rectangle {
        id: bg
        color: Style.bgColor
    }

    color: Style.textColor
    font.weight: Font.DemiBold
    font.pointSize: 10

    onCodeChanged: {
        switch(code) {
            case 0:
                text = "0: OK"
                bg.color = Style.greenColor
                break
            case 1:
                text = "1: CANCELLED"
                bg.color = Style.orangeColor
                break
            case 2:
                text = "2: UNKNOWN"
                bg.color = Style.redColor
                break
            case 3:
                text = "3: INVALID_ARGUMENT"
                bg.color = Style.orangeColor
                break
            case 4:
                text = "4: DEADLINE_EXCEEDED"
                bg.color = Style.orangeColor
                break
            case 5:
                text = "5: NOT_FOUND"
                bg.color = Style.orangeColor
                break
            case 6:
                text = "6: ALREADY_EXISTS"
                bg.color = Style.orangeColor
                break
            case 7:
                text = "7: PERMISSION_DENIED"
                bg.color = Style.redColor
                break
            case 8:
                text = "8: RESOURCE_EXHAUSTED"
                bg.color = Style.redColor
                break
            case 9:
                text = "9: FAILED_PRECONDITION"
                bg.color = Style.redColor
                break
            case 10:
                text = "10: ABORTED"
                bg.color = Style.redColor
                break
            case 11:
                text = "11: OUT_OF_RANGE"
                bg.color = Style.redColor
                break
            case 12:
                text = "12: UNIMPLEMENTED"
                bg.color = Style.orangeColor
                break
            case 13:
                text = "13: INTERNAL"
                bg.color = Style.redColor
                break
            case 14:
                text = "14: UNAVAILABLE"
                bg.color = Style.redColor
                break
            case 15:
                text = "15: DATA_LOSS"
                bg.color = Style.redColor
                break
            case 16:
                text = "16: UNAUTHENTICATED"
                bg.color = Style.redColor
                break
            default:
                text = ""
                bg.color = Style.bgColor
        }
    }
}
