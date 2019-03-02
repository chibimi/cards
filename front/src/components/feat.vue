<template>
	<div class="w-100">
		<div class="row px-3">
			<label class="col-1 col-form-label">VO</label>
			<input v-model="feat.original_name" type="text" class="form-control col-4" placeholder="English Name">
			<label class="col-1 col-form-label">Name</label>
			<input v-model="feat.name" type="text" class="form-control col-4" placeholder="French Name">
			<div class="col-2 pr-0">
				<button v-if="feat.id" type="submit" class="form-control btn btn-success" @click="save(feat)">Update</button>
				<button v-if="!feat.id" type="submit" class="form-control btn btn-primary" @click="save(feat)">Add</button>
			</div>
		</div>

		<div class="row px-3 mt-2">
			<textarea v-model="feat.description" type="text" class="form-control col" rows="3" placeholder/>
			
		</div>
	</div>
</template>

<script>
export default {
	name: "Feat",
	props: ["feat"],
	methods: {
		save: function(feat) {
			if (feat.id == null) {
				feat.id = 0;
			}
			this.$http
				.put("http://localhost:9901/feats/" + feat.id, feat)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						feat.id = res.data;
					}
				});
		},
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
