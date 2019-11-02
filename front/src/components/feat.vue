<template>
	<div class="w-100  mt-4" >

		<input v-model="feat.name" type="text" class="form-control" placeholder="Name">
		<textarea v-model="feat.description" type="text" class="form-control mt-2" rows="4" placeholder="Feat description"/>
		<textarea v-model="feat.fluff" type="text" class="form-control mt-2" rows="4" placeholder="Feat fluff (not mandatory)"/>
		<button type="submit" class="form-control btn btn-primary mt-2" @click="save(feat)">Save</button>
		


		<!-- <div class="row px-3">
			<label class="col-1 col-form-label">Name</label>
			<label class="col-1 col-form-label">VO</label>
			<input v-model="feat.original_name" type="text" class="form-control col-4" placeholder="Name">
			<label class="col-1 col-form-label">VF</label>
			<input v-model="feat.name" type="text" class="form-control col-4" placeholder="French Name">
			<div class="col-1 pr-0">
				<button v-if="feat.id" type="submit" class="form-control btn btn-success" @click="save(feat)">Update</button>
				<button v-if="!feat.id" type="submit" class="form-control btn btn-primary" @click="save(feat)">Add</button>
			</div>
		</div>
		<div class="row px-3 mt-2">
			<textarea v-model="feat.description" type="text" class="form-control col" rows="3" placeholder/>
		</div> -->
	</div>
</template>

<script>
export default {
	name: "Feat",
	props: ["ref_id"],
	watch: {
		ref_id: function(newVal) {
			this.get(newVal);
		}
	},
	created: function() {
		this.get(this.ref_id);
	},
	data() {
		return {
			feat: {
				ref_id: this.ref_id
			},
			update: false
		};
	},
	methods: {
		get: function(refID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/feat?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.feat = res.data;
					if (!this.feat.ref_id) {
						this.update = true;
					}
				});
		},
		save: function(feat) {
			if (feat.ref_id == null) {
				return
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + feat.ref_id + "/feat?lang=" + this.$language, feat)
				.then(function(res) {
					console.log(res);
					// if (res.feat === 201) {
					// 	feat.id = res.data;
					// 	this.update=false;
					// }
				});
		}
	}
};
</script>

<style scoped>
</style>
