<template>
	<div v-if="refs.length != 0">
		
		<div
			class="modal fade"
			id="exampleModal"
			tabindex="-1"
			role="dialog"
			aria-labelledby="exampleModalLabel"
			aria-hidden="true"
		>
			<div class="modal-dialog modal-dialog-centered" role="document">
				<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title" id="exampleModalLabel">{{ ref.title }}</h5>
						<button type="button" class="close" data-dismiss="modal" aria-label="Close">
							<span aria-hidden="true">&times;</span>
						</button>
					</div>
					<div class="modal-body">
						<h5>What's wrong ?</h5>
						<textarea v-model="feedback" placeholder=""></textarea>
						<input v-model="reviewer" class="mt-1" placeholder="reviewer name (optionnal)" />
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-primary" data-dismiss="modal" @click="saveFeedback(ref.id, index, 'bad', feedback, reviewer)">Save</button>
					</div>
				</div>
			</div>
		</div>

		<h4>{{ category.name }}</h4>
		<div class="d-flex flex-wrap">
			<div class="col-6" v-for="(r, i) in refs" :key="r.id">
				<div class="row" v-bind:class="{ odd: i % 4 in [0, 1] }">
					<div class="col-6 align-self-center">{{ r.title }} </div>

					<div class="col-6 align-self-center px-4">
						<div class="counter float-right ml-3" style="user-select: none">
							<span class="bg-success name_left" v-on:click="saveFeedback(r.id, i, 'good', '', '')"
								><i class="fas fa-thumbs-up"></i
							></span>
							<span class="count">{{ r.review_good || 0 }} | {{ r.review_bad || 0 }} </span>
							<span
								class="bg-danger name_right"
								data-toggle="modal"
								data-target="#exampleModal"
								v-on:click="ref = r; index = i; feedback= ''"
								><i class="fas fa-thumbs-down"></i
							></span>
						</div>

						<div class="counter float-right" style="user-select: none">
							<span class="bg-primary name_left" v-on:click="$emit('add', r.id)"
								><i class="fas fa-plus" style="user-select: none" href="#"></i
							></span>
							<span class="count">{{ card_ids[r.id] || 0 }}</span>
							<span class="bg-primary name_right" v-on:click="$emit('remove', r.id)"
								><i class="fas fa-minus" style="user-select: none" href="#"></i
							></span>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'Refs',
	props: ['card_ids', 'faction', 'category'],
	components: {},

	watch: {
		faction: function (newVal) {
			this.getRefs(newVal, this.category.id)
		},
		category: function (newVal) {
			this.getRefs(this.faction, newVal.id)
		},
	},
	created: function () {
		console.log('created', this.faction, this.category)
		this.getRefs(this.faction, this.category.id)
	},
	data() {
		return {
			refs: [],
			ref: {},
			index: null,
			feedback: "",
			reviewer: "",
		}
	},
	methods: {
		getRefs: function (faction, category) {
			if (!faction || !category) {
				return
			}
			this.$http
				.get(
					process.env.VUE_APP_API_ENDPOINT +
						`/ref?faction_id=${faction}&category_id=${category}&lang=fr&status=done`
				)
				.then(function (res) {
					console.debug(res)
					this.refs = res.data
				})
				.catch(function (err) {
					console.error(err)
				})
		},
		saveFeedback: function(ref_id, index, rating, feedback, reviewer){
			var review = {
				ref_id: ref_id,
				lang: 'fr',
				rating: rating,
				comment: feedback,
				reviewer: reviewer,
			}
			this.$http
				.post(process.env.VUE_APP_API_ENDPOINT + `/reviews`, review)
				.then(function (res) {
					console.debug(res)
					this.$http
						.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${ref_id}/rating?lang=fr`)
						.then(function (res){
							this.refs[index].review_good = res.data.good
							this.refs[index].review_bad = res.data.bad
						})
						.catch(function (err){
							console.error(err)
						})
				})
				.catch(function (err) {
					console.error(err)
				})
		}
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.counter {
	height: 20px;
	font-size: 14px;
	.name_left {
		display: inline-block;
		// float: left;
		padding: 0 6px 0 6px;
		color: white;
		border: solid 1px #bbbbbb;
		border-right: none;
		border-radius: 4px 0 0 4px;
	}
	.name_right {
		display: inline-block;
		// float: left;
		padding: 0 6px 0 6px;
		color: white;
		border: solid 1px #bbbbbb;
		border-right: none;
		border-left: none;
		border-radius: 0 4px 4px 0;
	}
	.count {
		position: relative;
		display: inline-block;
		// float: left;
		margin: 0;
		padding: 0 4px;
		border: solid 1px #dbdbdb;
		// border-radius: 0 4px 4px 0;
		background-color: white;
	}
}

.row {
	height: 60px;
}
.odd {
	background-color: #ececec;
}

.ref {
	.header {
		@extend .row;
		@extend .form-inline;
	}

	.error {
		@extend .row;
		@extend .alert;
		@extend .mx-0;
		@extend .py-1;
	}

	.nav {
		@extend .nav;
		@extend .nav-tabs;

		a {
			@extend .nav-item;
			@extend .nav-link;
		}
	}

	.content {
		// @extend .tab-content;
		@extend .container;
		@extend .px-0;
		@extend .mt-4;
	}
}
</style>
