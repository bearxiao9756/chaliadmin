
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="用户类型:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入用户类型" />
</el-form-item>
        <el-form-item label="访问哈希:" prop="accessHash">
    <el-input v-model.number="formData.accessHash" :clearable="true" placeholder="请输入访问哈希" />
</el-form-item>
        <el-form-item label="密钥 ID:" prop="secretKeyId">
    <el-input v-model.number="formData.secretKeyId" :clearable="true" placeholder="请输入密钥 ID" />
</el-form-item>
        <el-form-item label="名:" prop="firstName">
    <el-input v-model="formData.firstName" :clearable="true" placeholder="请输入名" />
</el-form-item>
        <el-form-item label="性:" prop="lastName">
    <el-input v-model="formData.lastName" :clearable="true" placeholder="请输入性" />
</el-form-item>
        <el-form-item label="用户名:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户名" />
</el-form-item>
        <el-form-item label="电话:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入电话" />
</el-form-item>
        <el-form-item label="国家代码:" prop="countryCode">
    <el-input v-model="formData.countryCode" :clearable="true" placeholder="请输入国家代码" />
</el-form-item>
        <el-form-item label="是否已验证:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="支持是否:" prop="support">
    <el-switch v-model="formData.support" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否欺诈:" prop="scam">
    <el-switch v-model="formData.scam" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否伪造:" prop="fake">
    <el-switch v-model="formData.fake" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否特权:" prop="premium">
    <el-switch v-model="formData.premium" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="特权到期时间:" prop="premiumExpireDate">
    <el-input v-model.number="formData.premiumExpireDate" :clearable="true" placeholder="请输入特权到期时间" />
</el-form-item>
        <el-form-item label="用户签名:" prop="about">
    <el-input v-model="formData.about" :clearable="true" placeholder="请输入用户签名" />
</el-form-item>
        <el-form-item label="用户状态:" prop="state">
    <el-input v-model.number="formData.state" :clearable="true" placeholder="请输入用户状态" />
</el-form-item>
        <el-form-item label="是否机器人:" prop="isBot">
    <el-switch v-model="formData.isBot" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="token 有效期:" prop="accountDaysTtl">
    <el-input v-model.number="formData.accountDaysTtl" :clearable="true" placeholder="请输入token 有效期" />
</el-form-item>
        <el-form-item label="照片 ID:" prop="photoId">
    <el-input v-model.number="formData.photoId" :clearable="true" placeholder="请输入照片 ID" />
</el-form-item>
        <el-form-item label="是否受限制的:" prop="restricted">
    <el-switch v-model="formData.restricted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="限制原因字段:" prop="restrictionReason">
    <el-input v-model="formData.restrictionReason" :clearable="true" placeholder="请输入限制原因字段" />
</el-form-item>
        <el-form-item label="存档并静音新的非联系同伴:" prop="archiveAndMuteNewNoncontactPeers">
    <el-switch v-model="formData.archiveAndMuteNewNoncontactPeers" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="emojiStatusDocumentID:" prop="emojiStatusDocumentId">
    <el-input v-model.number="formData.emojiStatusDocumentId" :clearable="true" placeholder="请输入emojiStatusDocumentID" />
</el-form-item>
        <el-form-item label="emojiStatusUntil:" prop="emojiStatusUntil">
    <el-input v-model.number="formData.emojiStatusUntil" :clearable="true" placeholder="请输入emojiStatusUntil" />
</el-form-item>
        <el-form-item label="storiesMaxId:" prop="storiesMaxId">
    <el-input v-model.number="formData.storiesMaxId" :clearable="true" placeholder="请输入storiesMaxId" />
</el-form-item>
        <el-form-item label="颜色:" prop="color">
    <el-input v-model.number="formData.color" :clearable="true" placeholder="请输入颜色" />
</el-form-item>
        <el-form-item label="colorBackgroundEmojiId字段:" prop="colorBackgroundEmojiId">
    <el-input v-model.number="formData.colorBackgroundEmojiId" :clearable="true" placeholder="请输入colorBackgroundEmojiId字段" />
</el-form-item>
        <el-form-item label="profileColor:" prop="profileColor">
    <el-input v-model.number="formData.profileColor" :clearable="true" placeholder="请输入profileColor" />
</el-form-item>
        <el-form-item label="profileColorBackgroundEmojiIo:" prop="profileColorBackgroundEmojiId">
    <el-input v-model.number="formData.profileColorBackgroundEmojiId" :clearable="true" placeholder="请输入profileColorBackgroundEmojiIo" />
</el-form-item>
        <el-form-item label="生日:" prop="birthday">
    <el-input v-model="formData.birthday" :clearable="true" placeholder="请输入生日" />
</el-form-item>
        <el-form-item label="个人频道ID:" prop="personalChannelId">
    <el-input v-model.number="formData.personalChannelId" :clearable="true" placeholder="请输入个人频道ID" />
</el-form-item>
        <el-form-item label="授权有效期天数:" prop="authorizationTtlDays">
    <el-input v-model.number="formData.authorizationTtlDays" :clearable="true" placeholder="请输入授权有效期天数" />
</el-form-item>
        <el-form-item label="是否删除:" prop="deleted">
    <el-switch v-model="formData.deleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="删除原因:" prop="deleteReason">
    <el-input v-model="formData.deleteReason" :clearable="true" placeholder="请输入删除原因" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新更新:" prop="updatedAt">
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
  createTeamGramUsers,
  updateTeamGramUsers,
  findTeamGramUsers
} from '@/api/system/teamgramusers'

defineOptions({
    name: 'TeamGramUsersForm'
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
            userType: undefined,
            accessHash: undefined,
            secretKeyId: undefined,
            firstName: '',
            lastName: '',
            username: '',
            phone: '',
            countryCode: '',
            verified: false,
            support: false,
            scam: false,
            fake: false,
            premium: false,
            premiumExpireDate: undefined,
            about: '',
            state: undefined,
            isBot: false,
            accountDaysTtl: undefined,
            photoId: undefined,
            restricted: false,
            restrictionReason: '',
            archiveAndMuteNewNoncontactPeers: false,
            emojiStatusDocumentId: undefined,
            emojiStatusUntil: undefined,
            storiesMaxId: undefined,
            color: undefined,
            colorBackgroundEmojiId: undefined,
            profileColor: undefined,
            profileColorBackgroundEmojiId: undefined,
            birthday: '',
            personalChannelId: undefined,
            authorizationTtlDays: undefined,
            deleted: false,
            deleteReason: '',
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
      const res = await findTeamGramUsers({ ID: route.query.id })
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
               res = await createTeamGramUsers(formData.value)
               break
             case 'update':
               res = await updateTeamGramUsers(formData.value)
               break
             default:
               res = await createTeamGramUsers(formData.value)
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
