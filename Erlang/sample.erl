-module (sample).
-export ([perm_sum/1]).

perm_sum(1) ->[[1]];
perm_sum(N) when N > 1-> 
	Head = [N],
	Seq=[{S, lists:filter(fun(X)-> lists:max(X) =< S end ,perm_sum(N-S))} || S <-lists:reverse(lists:seq(1,N-1))],
	Result=lists:foldr(fun ({Key,PermsList}, Acc1) -> lists:foldr(fun(P, Acc)-> [[Key|P]|Acc] end, Acc1, PermsList) end, [], Seq),
	[Head|Result].
