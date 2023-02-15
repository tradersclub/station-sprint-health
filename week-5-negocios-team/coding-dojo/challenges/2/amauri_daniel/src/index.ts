
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


class AdaptadorEasyImage implements Image {
    async resize(inputPath: string, outputPath: string, width: number, height: number): Promise<void> {
        await easyimageResize({
            src: inputPath,
            width: width,
            height: height,
            dst: outputPath,
            })
    }
}

const easyWhatever = new AdaptadorEasyImage();

const something = easyWhatever.resize(imagePath, imagePathDestination, 10, 10);
// console.log(sum(2,3))
console.log('hello world');