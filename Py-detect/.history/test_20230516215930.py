import numpy as np
from torch import load, from_numpy, tensor
# import torch
from utils.general import (LOGGER, check_file, check_img_size, check_imshow, check_requirements, colorstr,
                        increment_path, non_max_suppression, print_args, scale_segments, strip_optimizer, xyxy2xywh)
from utils.augmentations import letterbox
import cv2
import torch

# 初始化参数
# 初始化目录
# FILE = Path(__file__).resolve()
# ROOT = FILE.parents[0]  # 定义YOLOv5的根目录
# if str(ROOT) not in sys.path:
#     sys.path.append(str(ROOT))  # 将YOLOv5的根目录添加到环境变量中（程序结束后删除）
# ROOT = Path(os.path.relpath(ROOT, Path.cwd()))
#
# weights = ROOT / 'yolov5s.pt'  # 权重文件地址   .pt文件
# source = ROOT / 'data/images'  # 测试数据文件(图片或视频)的保存路径
# data = ROOT / 'data/coco128.yaml'  # 标签文件地址   .yaml文件


imgsz = (200, 200)  # 输入图片的大小 默认640(pixels)
conf_thres = 0.25  # object置信度阈值 默认0.25  用在nms中
iou_thres = 0.45  # 做nms的iou阈值 默认0.45   用在nms中
max_det = 1000  # 每张图片最多的目标数量  用在nms中
device = 'cpu'  # 设置代码执行的设备 cuda device, i.e. 0 or 0,1,2,3 or cpu
classes = None  # 在nms中是否是只保留某些特定的类 默认是None 就是所有类只要满足条件都可以保留 --class 0, or --class 0 2 3
agnostic_nms = False  # 进行nms是否也除去不同类别之间的框 默认False
augment = False  # 预测是否也要采用数据增强 TTA 默认False
visualize = False  # 特征图可视化 默认FALSE
half = False  # 是否使用半精度 Float16 推理 可以缩短推理时间 但是默认是False
dnn = False  # 使用OpenCV DNN进行ONNX推理

# device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
# device = select_device(device)
# print(device)


# 导入模型

model = load('./weights/best.pt')["model"].float().eval()  # load FP32 model
imgsz = check_img_size(imgsz, s=model.stride.max())  # check img_size
if half:
    model.half()  # to FP16

# Get names and colors
names = model.module.names if hasattr(model, 'module') else model.names  # 获取标签


# colors = [[random.randint(0, 255) for _ in range(3)] for _ in names]


def detect_(img):
    im0 = img
    im = letterbox(im0)[0]
    im = im.transpose((2, 0, 1))[::-1]  # HWC to CHW, BGR to RGB
    im = np.ascontiguousarray(im)  # www 函数将一个内存不连续存储的数组转换为内存连续存储的数组，使得运行速度更快。
    im = from_numpy(im).to(device)
    im = im.half() if half else im.float()  # uint8 to fp16/32
    im /= 255  # 0 - 255 to 0.0 - 1.0
    if im.ndimension() == 3:
        im = im[None]  # 增加维度

    # Inference
    pred = model(im)[0]
    # NMS
    pred = non_max_suppression(pred, conf_thres, iou_thres, classes, agnostic_nms, max_det=max_det)

    # 用于存放结果
    detections = []
    gn = torch.tensor(im0.shape)[[1, 0, 1, 0]]
    # gn = torch.tensor([200,200,200,200])
    print(gn)
    for i, det in enumerate(pred):  # per image 每张图片
        if len(det):
            # det[:, :4] = scale_segments(im.shape[2:], det[:, :4], im0.shape).round()
            # result
            for *xyxy, conf, cls in reversed(det):
                xywh = (xyxy2xywh(torch.tensor(xyxy).view(1, 4))).view(-1).tolist()
                # xyxy = (xywhn2xyxy(torch.tensor(xyxy).view(1, 4), w=200, h=200, padw=0, eps=0)).view(-1).tolist()
                # xywh = (xyxy2xywhn(torch.tensor(xyxy).view(1, 4), w=200, h=200, clip=False, eps=0.0)).view(-1).tolist()
                # xywh = [round(x) for x in xywh]
                # xywh = [xywh[0] - xywh[2] // 2, xywh[1] - xywh[3] // 2, xywh[2],
                #         xywh[3]]  # 检测到目标位置，格式：（left，top，w，h）
                                    # [x_min,y_min,w,h] -> [x_min,y_min,x_max,y_max] -> [x_min,y_min,w,h]
                # xywh = [x_center, y_center, width, height] convert to [x_min,y_min,x_max,y_max]

                # x_min = xywh[0] - xywh[2] // 2
                # y_min = xywh[1] + xywh[3] // 2
                # x_max = xywh[0] + xywh[2] // 2
                # y_max = xywh[1] - xywh[3] // 2


                # x_min = xywh[0]-xywh[2]//2
                # y_min = xywh[1]-xywh[3]//2
                # x_max = xywh[0]+xywh[2]//2
                # y_max = xywh[1]+xywh[3]//2
                # xywh = [x_min,y_min,x_max,y_max]

                cls = names[int(cls)]
                conf = float(conf)
                detections.append({'class': cls, 'conf': conf, 'position': xywh})

    #     # 输出结果
    # for i in detections:
    #     print(i)
    return detections


if __name__ == '__main__':
    pic = cv2.imread('./img/inclusion_205.jpg')
    # pic = cv2.imread('./img/scratches_287.jpg')
    data = detect_(pic)
    for i in data:
        print(data)
