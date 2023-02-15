import {sum} from "./utils";

import { resize as easyimageResize } from "easyimage";

// sempre jpg
interface Image {
    resize(
        inputPath: string,
        outputPath: string,
         width: number, height: number
        ): void;
}

const imagePath = './imagens/cachorro.jpeg'
const imagePathDestination = './imagens/resized'

// ;(async () => {
//     await easyimageResize({
//     src: imagePath,
//     width: 50,
//     height: 50,
//     dst: imagePathDestination,
//     })
// })();


class RedimensionaImagem implements Image {
    async resize(inputPath: string, outputPath: string, width: number, height: number): void {
        await easyimageResize({
            src: inputPath,
            width: width,
            height: height,
            dst: outputPath,
            })
    }
}

// console.log(sum(2,3))
