

import { resize as easyimageResize } from "easyimage";

interface Image {
    resize(
        inputPath: string,
        outputPath: string,
         width: number, height: number
        ): Promise<void>;
}
class EasyImageAdapter implements Image {
    async resize(inputPath: string, outputPath: string, width: number, height: number): Promise<void> {
        await easyimageResize({
            src: inputPath,
            width: width,
            height: height,
            dst: outputPath,
            })
    }
}

class ImageTool {
    private imageResizerAdapter: Image;

    constructor(imageResizerAdapter: Image) {
        this.imageResizerAdapter = imageResizerAdapter
    }

    resizer(inputPath: string, outputPath: string, width: number, height: number) {
        return this.imageResizerAdapter.resize(inputPath, outputPath, width, height)
    }
}   


const imagePath = `${__dirname}/../imagens/cachorro.jpg`
const imagePathDestination = `${__dirname}/../imagens/resized/cachorro40x40${new Date().getTime()}.jpg`



;(async () => {
    const easy = new ImageTool(new EasyImageAdapter())
    await easy.resizer(imagePath, imagePathDestination, 40, 40)  
})();
