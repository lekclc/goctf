<template>
    <div>
        <h1>Game Page</h1>
        <button @click="AddChallenge">添加题目</button>
    </div>
</template>

<script>
export default {
    name: 'Game',
    return: {
        games: []
    },
    async created() {
        //从url中获取参数
        const id = this.$route.params.id
        console.log(id)
    },
    methods: {
        AddChallenge() {
            this.$router.push('/AddChallenge' )
        }
    },
}

</script>