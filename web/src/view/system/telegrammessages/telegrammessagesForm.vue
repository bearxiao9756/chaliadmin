
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="用户ID:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户ID" />
</el-form-item>
        <el-form-item label="消盒子序号:" prop="userMessageBoxId">
    <el-input v-model.number="formData.userMessageBoxId" :clearable="true" placeholder="请输入消盒子序号" />
</el-form-item>
        <el-form-item label="发送者会话编号:" prop="dialogId1">
    <el-input v-model.number="formData.dialogId1" :clearable="true" placeholder="请输入发送者会话编号" />
</el-form-item>
        <el-form-item label="接收者会话编号:" prop="dialogId2">
    <el-input v-model.number="formData.dialogId2" :clearable="true" placeholder="请输入接收者会话编号" />
</el-form-item>
        <el-form-item label="消息ID:" prop="dialogMessageId">
    <el-input v-model.number="formData.dialogMessageId" :clearable="true" placeholder="请输入消息ID" />
</el-form-item>
        <el-form-item label="发送者ID:" prop="senderUserId">
    <el-input v-model.number="formData.senderUserId" :clearable="true" placeholder="请输入发送者ID" />
</el-form-item>
        <el-form-item label="消息类型:" prop="peerType">
    <el-input v-model.number="formData.peerType" :clearable="true" placeholder="请输入消息类型" />
</el-form-item>
        <el-form-item label="消息类型ID:" prop="peerId">
    <el-input v-model.number="formData.peerId" :clearable="true" placeholder="请输入消息类型ID" />
</el-form-item>
        <el-form-item label="幂等随机数:" prop="randomId">
    <el-input v-model.number="formData.randomId" :clearable="true" placeholder="请输入幂等随机数" />
</el-form-item>
        <el-form-item label="消息过滤类型:" prop="messageFilterType">
    <el-input v-model.number="formData.messageFilterType" :clearable="true" placeholder="请输入消息过滤类型" />
</el-form-item>
        <el-form-item label="消息内容:" prop="message">
    <el-input v-model="formData.message" :clearable="true" placeholder="请输入消息内容" />
</el-form-item>
        <el-form-item label="是否提及:" prop="mentioned">
    <el-switch v-model="formData.mentioned" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="未读媒体:" prop="mediaUnread">
    <el-switch v-model="formData.mediaUnread" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否置顶:" prop="pinned">
    <el-switch v-model="formData.pinned" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否反应:" prop="hasReaction">
    <el-switch v-model="formData.hasReaction" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="反应内容:" prop="reaction">
    <el-input v-model="formData.reaction" :clearable="true" placeholder="请输入反应内容" />
</el-form-item>
        <el-form-item label="反应时间:" prop="reactionDate">
    <el-input v-model.number="formData.reactionDate" :clearable="true" placeholder="请输入反应时间" />
</el-form-item>
        <el-form-item label="未读反应:" prop="reactionUnread">
    <el-switch v-model="formData.reactionUnread" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="时间:" prop="date2">
    <el-input v-model.number="formData.date2" :clearable="true" placeholder="请输入时间" />
</el-form-item>
        <el-form-item label="过期时间:" prop="ttlPeriod">
    <el-input v-model.number="formData.ttlPeriod" :clearable="true" placeholder="请输入过期时间" />
</el-form-item>
        <el-form-item label="保存类型:" prop="savedPeerType">
    <el-input v-model.number="formData.savedPeerType" :clearable="true" placeholder="请输入保存类型" />
</el-form-item>
        <el-form-item label="保存类型ID:" prop="savedPeerId">
    <el-input v-model.number="formData.savedPeerId" :clearable="true" placeholder="请输入保存类型ID" />
</el-form-item>
        <el-form-item label="是否删除:" prop="deleted">
    <el-switch v-model="formData.deleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMessages,
  updateMessages,
  findMessages
} from '@/api/system/telegrammessages'

defineOptions({
    name: 'MessagesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            userId: undefined,
            userMessageBoxId: undefined,
            dialogId1: undefined,
            dialogId2: undefined,
            dialogMessageId: undefined,
            senderUserId: undefined,
            peerType: undefined,
            peerId: undefined,
            randomId: undefined,
            messageFilterType: undefined,
            message: '',
            mentioned: false,
            mediaUnread: false,
            pinned: false,
            hasReaction: false,
            reaction: '',
            reactionDate: undefined,
            reactionUnread: false,
            date2: undefined,
            ttlPeriod: undefined,
            savedPeerType: undefined,
            savedPeerId: undefined,
            deleted: false,
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMessages({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createMessages(formData.value)
               break
             case 'update':
               res = await updateMessages(formData.value)
               break
             default:
               res = await createMessages(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
