<template>
    <div v-if="visible" class="modal-backdrop">
        <div class="modal-content">
            <p>{{ message }}</p>
            <button class="btn" @click="confirm">确认</button>
            <button class="btn" @click="cancel">取消</button>
        </div>
    </div>
</template>

<script>
export default {
    props: ['message'],  // 接收外部传入的消息
    data() {
        return {
            visible: false
        };
    },
    methods: {
        confirm() {
            this.$emit('confirm');
            this.visible = false;
        },
        cancel() {
            this.visible = false;
        },
        show() {
            this.visible = true;
        }
    }
};
</script>

<style scoped>
.modal-backdrop {
    position: fixed;           /* 固定位置 */
    top: 0;                    /* 背景从顶部开始 */
    left: 0;                   /* 背景从左边开始 */
    width: 100%;               /* 背景覆盖整个宽度 */
    height: 100%;              /* 背景覆盖整个高度 */
    background: rgba(0, 0, 0, 0.5);  /* 半透明背景 */
    z-index: 9999;             /* 置于最顶层 */
}

.modal-content {
    position: fixed;           /* 固定模态框 */
    top: 50%;                  /* 垂直居中 */
    left: 50%;                 /* 水平居中 */
    transform: translate(-50%, -50%);  /* 移动模态框到正确的居中位置 */
    background: #2D3047;
    padding: 20px;
    border-radius: 10px;
    width: 300px;              /* 根据需要设置宽度 */
    text-align: center;
    z-index: 10000;            /* 比 backdrop 高一层 */
    border: 3px solid #F3CA40;
    color:#FFFFFF;
    font-size: 20px;
}


.btn {
    margin-top: 30px;
    margin-left: 11px;
    width: 90px;
    height: 40px;
    background-color: #9C89B8;
    border: 3px solid #1C1C1E;
    color: #fff;
    border-radius: 4px;
    cursor: pointer;
    font-size: 15px;
}
.btn:hover{
    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
}

</style>
