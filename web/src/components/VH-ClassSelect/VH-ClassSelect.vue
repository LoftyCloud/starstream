<template>
    <div class="class-box">
        <div class="class-inner-box">
            <!-- "全部" 按钮 -->
            <button :class="select == -1 ? 'btn avatice' : 'btn'" @click="onclick(0, -1)">全部</button>

            <!-- 分类按钮 -->
            <button
                :class="select == key ? 'btn avatice' : 'btn'"
                v-for="(item, key) in listData"
                :key="key"
                @click="onclick(item, key)">
                {{ item["item"] }}
            </button>
        </div>
    </div>
</template>



<script>
export default {
    name: 'VH-ClassSelect',
    data () {
        return {
            listData:[
                {m: '1', item: '可爱', key: 1},
                {m: '2', item: '搞笑', key: 2},
                {m: '3', item: '学习', key: 3},
                {m: '4', item: '游戏', key: 4},
                {m: '5', item: '日常', key: 5},
                {m: '6', item: '正经', key: 6},
                {m: '7', item: '其他', key: 7}
            ],
            select:-1,
        }
    },

    props:{
        itemData: {
            type: Array,
            default: () => ([
                { m: '1', item: '可爱', key: 1 },
                { m: '2', item: '搞笑', key: 2 },
                { m: '3', item: '学习', key: 3 },
                { m: '4', item: '游戏', key: 4 },
                { m: '5', item: '日常', key: 5 },
                { m: '6', item: '正经', key: 6 },
                { m: '7', item: '其他', key: 7 }
            ])
        },
        // label: {
        //     type: String,
        //     default: 'item' // 设置默认值为 'item'
        // },
        vhKey: {
            type: String,
            default: 'key' // 设置默认的 key 值
        },
        cid: {
            type: Number,
            default: 0
        }
    },

    watch: {
        itemData(news) {
            if (news && news.length > 0) {
                this.listData = news;
            }
        },
        cid(news) {
            if (news === 0) {
                this.select = -1; // 重置为 "全部" 被选中
            }
        }
    },

    created(){
		// this.listData = this.itemData;
    },
    methods:{
        onclick(item,keys){
            this.select=keys

            if(item == 0){  // 搜索全部视频
                this.$emit('select',0)
            }else{  // 分类搜索
                let key = item["key"]
                // console.log(keys)
                this.$emit('select',key)
            }
        }
    },
}
</script>

<style  scoped>
*{
    margin: 0;
    padding: 0;
}
.class-box{
    width: 100%;
    height: 65px;
    box-sizing: content-box;
    padding: 0 38px;
    position: relative;
    
}
.class-inner-box{
    width: 1260px;
    height: 65px;
    background-color: #2D3047;
    box-shadow: 0 3px 12px #2D3047;
    color: #fff;
    box-sizing: content-box;
    border-radius: 6px;
    position: absolute;
    overflow: hidden;
    text-align: center;
    z-index: 1;
}
.btn{
    min-width: 90px;
    display: inline-block;
    box-sizing: content-box;
    width: auto;
    height: 32px;
    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
    border-radius: 6px;
    text-align: center;
    font-weight: 400;
    transition: background-color .3s,color .3s;
    cursor: pointer;
    padding: 5px 10px;
    margin: 10px;
    border: 3px solid #1C1C1E;
    font-size: 16px;
}

.btn:hover{
    background-color: #9C89B8;
    color: #FFFFFF;
}

.class-inner-box:hover{
    box-shadow:  -3px -3px 9px #F3CA40, 3px 3px 7px #9C89B8;
}

@keyframes showDropUp{
	0%{
		transform:translateY(-5px);
        height: 65px;
	}
	100%{
		transform:translateY(0px);
        height: 256px;
        overflow-y: auto;
	}
}
.avatice{
    font-size: 20px;
    color: #FFFFFF;
    background-color: #9C89B8;
}
</style>
