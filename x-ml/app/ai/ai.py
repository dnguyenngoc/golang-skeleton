from ai.face_detection import FaceDetection
from ai.face_recognition import FaceRecognition
from settings import ml_config
import json


class CompleteModel:
    def __init__(self):
        self.face_detection, self.face_recognition, self.classes = self._load_model()
    
    @staticmethod
    def _load_model():
        face_detection = FaceDetection()
        face_detection.prepare(ctx_id=0, det_size=(640, 640))
        face_recognition = FaceRecognition()
        f = open(ml_config.FACE_REC_LABLE_PATH, encoding="utf8")
        classes = json.loads(f.read())
        classes_new = {}
        for key, value in classes.items():
            classes_new[value] = key
        return face_detection, face_recognition, classes_new
    